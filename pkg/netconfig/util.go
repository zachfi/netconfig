package netconfig

import (
	"bytes"
	"io/ioutil"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/xaque208/netconfig/pkg/netconfig/data"
	yaml "gopkg.in/yaml.v2"
)

// loadHostDataFile unmarshals a YAML file into the received interface{} or returns an error.
func loadHostDataFile(filename string, d *data.HostData) error {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.UnmarshalStrict(yamlFile, d)
	if err != nil {
		return err
	}

	return nil
}

// loadDataConfig unmarshals a YAML file into the received interface{} or returns an error.
func loadDataConfig(filename string, d *data.Data) error {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.UnmarshalStrict(yamlFile, d)
	if err != nil {
		return err
	}

	return nil
}

// templateStringsForDevice renders a list of template strings given an inventory.NetworkHost object.
func templateStringsForDevice(host Host, templates []string) []string {
	var strings []string

	for _, t := range templates {
		tmpl, err := template.New("template").Parse(t)
		if err != nil {
			log.Error(err)
		}

		var buf bytes.Buffer

		err = tmpl.Execute(&buf, host)
		if err != nil {
			log.Error(err)
		}

		strings = append(strings, buf.String())
	}

	return strings
}
