package netconfig

type Config struct {
	Junos        JunosConfig `yaml:"junos,omitempty"`
	OtelEndpoint string      `yaml:"otel_endpoint"`
}

// JunosConfig is the configuration for Junos devices.
type JunosConfig struct {
	Hosts      []string `yaml:"hosts,omitempty"`
	Username   string   `yaml:"username,omitempty"`
	PrivateKey string   `yaml:"private_key,omitempty"`
}
