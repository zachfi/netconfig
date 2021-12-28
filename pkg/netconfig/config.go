package netconfig

import "github.com/xaque208/znet/modules/inventory"

type Config struct {
	Junos        JunosConfig      `yaml:"junos,omitempty"`
	OtelEndpoint string           `yaml:"otel_endpoint"`
	Data         DataConfig       `yaml:"data"`
	Inventory    inventory.Config `yaml:"inventory"`
}

// JunosConfig is the configuration for Junos devices.
type JunosConfig struct {
	Hosts      []string `yaml:"hosts,omitempty"`
	Username   string   `yaml:"username,omitempty"`
	PrivateKey string   `yaml:"private_key,omitempty"`
}

// DataConfig is the configuration for data.
type DataConfig struct {
	Directory string `yaml:"directory,omitempty"`
}
