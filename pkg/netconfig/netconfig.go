package netconfig

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/imdario/mergo"
	"github.com/scottdware/go-junos"
	"google.golang.org/protobuf/proto"

	"github.com/xaque208/znet/modules/inventory"
)

// Host is a single configurable host.
type Host struct {
	HostName    string
	NetworkHost *inventory.NetworkHost
	Data        HostData
	Environment map[string]string
}

// NetConfig is enough data to configure some network hosts.
type NetConfig struct {
	logger log.Logger
	cfg    *Config

	junosAuth *junos.AuthMethod
	ldap      *inventory.LDAPInventory

	Data  Data
	Hosts []Host
}

// NewNetConfig is used to build a new *NetConfig.
func NewNetConfig(cfg Config, logger log.Logger) (*NetConfig, error) {
	logger = log.With(logger, "module", "timer")
	n := &NetConfig{
		logger: logger,
		cfg:    &cfg,
		junosAuth: &junos.AuthMethod{
			Username:   cfg.Junos.Username,
			PrivateKey: cfg.Junos.PrivateKey,
		},
	}

	data, err := loadData(cfg.Data.Directory, logger)
	if err != nil {
		return nil, err
	}
	n.Data = data

	inv, err := inventory.NewLDAPInventory(cfg.Inventory, logger)
	if err != nil {
		return nil, err
	}
	n.ldap = inv

	hosts, err := inv.ListNetworkHosts(context.TODO())
	if err != nil {
		return nil, err
	}

	_ = level.Debug(logger).Log("msg", "netconfig", "host_count", len(hosts))

	for i, h := range hosts {
		netHost := proto.Clone(&hosts[i])

		host := Host{
			NetworkHost: netHost.(*inventory.NetworkHost),
			HostName:    strings.Join([]string{h.Name, h.Domain}, "."),
			// Environment: env,
		}

		d := n.dataForHost(host)
		host.Data = d

		n.Hosts = append(n.Hosts, host)
	}

	return n, nil
}

// LoadData receives a configuration directory from which to load the data for Znet.
func loadData(configDir string, logger log.Logger) (Data, error) {
	_ = level.Debug(logger).Log("msg", "loading data", "path", configDir)
	dataConfig := Data{}
	err := loadYamlFile(fmt.Sprintf("%s/%s", configDir, "data.yaml"), &dataConfig)
	if err != nil {
		return Data{}, fmt.Errorf("failed to load yaml file %s: %s", configDir, err)
	}

	return dataConfig, nil
}

// ConfigureNetwork configures all discovered network devices.
func (n *NetConfig) ConfigureNetwork(commit bool, confirm int, diff bool) error {
	if n == nil {
		return fmt.Errorf("unable to configure network with nil NetConfig")
	}

	wg := sync.WaitGroup{}
	for _, host := range n.Hosts {
		wg.Add(1)
		go func(h Host) {
			if h.NetworkHost.Platform == "junos" {
				_ = level.Debug(n.logger).Log("msg", "configuring", "host", h.HostName)

				err := n.ConfigureNetworkHost(h, commit, confirm, diff)
				if err != nil {
					_ = level.Error(n.logger).Log("msg", "failed to configure", "host", h.HostName, "err", err)
				}
			}

			wg.Done()
		}(host)
	}
	wg.Wait()

	return nil
}

// ConfigureNetworkHost renders the templates using associated data for a
// network host.  The hosts about which to load the templates, are retrieved
// from LDAP.
func (n *NetConfig) ConfigureNetworkHost(host Host, commit bool, confirm int, diff bool) error {
	session, err := junos.NewSession(host.HostName, n.junosAuth)
	if err != nil {
		return err
	}

	defer session.Close()

	templates := n.templatesForDevice(host)

	_ = level.Debug(n.logger).Log("msg", "templates for device", "count", len(templates))

	var renderedTemplates []string
	for _, t := range templates {
		result := n.renderHostTemplateFile(host, t)
		renderedTemplates = append(renderedTemplates, result)
	}

	if diff {
		_ = level.Debug(n.logger).Log("msg", "rendered templates", "output", renderedTemplates)
	}

	err = session.Lock()
	if err != nil {
		return fmt.Errorf("unable to lock session on %s: %s", host.HostName, err)
	}

	defer func() {
		err = session.Unlock()
		if err != nil {
			_ = level.Error(n.logger).Log("msg", "error unlocking session", "host", host.HostName, "err", err)
		}
	}()

	err = session.Config(renderedTemplates, "text", false)
	if err != nil {
		return fmt.Errorf("unable to load configuration on %s: %s", host.HostName, err)
	}

	diffResult, err := session.Diff(0)
	if err != nil {
		return err
	}

	if len(diffResult) > 1 {
		_ = level.Info(n.logger).Log("msg", "configuration changes", "host", host.HostName)
		fmt.Printf("%+v", diffResult)

		if commit {
			if confirm > 0 {
				err = session.CommitConfirm(confirm)
				if err != nil {
					return err
				}
			} else {
				err = session.Commit()
				if err != nil {
					return err
				}
			}

			err = session.Commit()
			if err != nil {
				return err
			}
		} else {
			err = session.Config("rollback", "text", false)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// DataForDevice returns HostData for a given NetworkHost.
func (n *NetConfig) dataForHost(host Host) HostData {
	hostData := HostData{}

	for _, f := range n.hierarchyForDevice(host) {
		fileHostData := HostData{}
		err := loadYamlFile(f, &fileHostData)
		if err != nil {
			_ = level.Error(n.logger).Log("msg", "failed to load yaml file", "err", err)
		}

		if err := mergo.Merge(&hostData, fileHostData, mergo.WithOverride); err != nil {
			_ = level.Error(n.logger).Log("msg", "failed to merge data", "err", err)
		}
	}

	return hostData
}

// HierarchyForDevice returns a list of file paths to consult for the data hierarchy.
func (n *NetConfig) hierarchyForDevice(host Host) []string {
	var files []string

	paths := templateStringsForDevice(host, n.Data.Hierarchy)

	for _, p := range paths {
		templateAbs := fmt.Sprintf("%s/%s", n.cfg.Data.Directory, p)
		if _, err := os.Stat(templateAbs); err == nil {
			files = append(files, templateAbs)
		}
	}

	return files
}

// templatesForDevice returns a list of template paths for a given host.
func (n *NetConfig) templatesForDevice(host Host) []string {
	var files []string

	_ = level.Debug(n.logger).Log("msg", "loading templates for host", "host", host.HostName)

	paths := templateStringsForDevice(host, n.Data.TemplatePaths)

	for _, p := range paths {
		templateAbs := fmt.Sprintf("%s/%s/%s", n.cfg.Data.Directory, n.Data.TemplateDir, p)
		if _, err := os.Stat(templateAbs); err == nil {
			globPattern := fmt.Sprintf("%s/*.tmpl", templateAbs)
			foundFiles, globErr := filepath.Glob(globPattern)
			if globErr != nil {
				_ = level.Error(n.logger).Log("msg", "failed to glob pattern", "err", globErr)
			} else {
				files = append(files, foundFiles...)
			}
		}
	}

	return files
}

// RenderHostTemplateFile renders a template file using a Host object.
func (n *NetConfig) renderHostTemplateFile(host Host, path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		_ = level.Error(n.logger).Log("msg", "failed to read path", "err", err)
	}

	str := string(b)
	tmpl, err := template.New("test").Parse(str)
	if err != nil {
		_ = level.Error(n.logger).Log("msg", "failed to parse template", "err", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, host)
	if err != nil {
		_ = level.Error(n.logger).Log("msg", "failed to execute template", "err", err)
	}

	return buf.String()
}
