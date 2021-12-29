package netconfig

import (
	"flag"

	"github.com/xaque208/znet/modules/inventory"
)

type Config struct {
	Junos           JunosConfig      `yaml:"junos,omitempty"`
	OtelEndpoint    string           `yaml:"otel_endpoint"`
	Data            DataConfig       `yaml:"data"`
	Inventory       inventory.Config `yaml:"inventory"`
	Commit          bool
	Diff            bool
	CommitConfirmed int
}

// JunosConfig is the configuration for Junos devices.
type JunosConfig struct {
	Hosts    []string `yaml:"hosts,omitempty"`
	Username string   `yaml:"username,omitempty"`
	Keyfile  string   `yaml:"keyfile,omitempty"`
}

// DataConfig is the configuration for data.
type DataConfig struct {
	Directory string `yaml:"directory,omitempty"`
}

func (c *Config) RegisterFlagsAndApplyDefaults(prefix string, f *flag.FlagSet) {
	c.CommitConfirmed = 0
	f.StringVar(&c.Junos.Username, "junos.username", "", "")
	f.StringVar(&c.Junos.Keyfile, "junos.keyfile", "", "")
	f.StringVar(&c.OtelEndpoint, "otel_endpoint", "", "otel endpoint, eg: tempo:4317")
	f.BoolVar(&c.Commit, "commit", false, "commit the diff")
	f.BoolVar(&c.Diff, "diff", true, "show the diff")
}
