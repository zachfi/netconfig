package data

type SecurityNAT struct {
	Source      SourceNAT      `yaml:"source"`
	Destination DestinationNAT `yaml:"destination"`
}

// SourceNAT is a set of NAT rules for SNAT.
type SourceNAT struct {
	RuleSets []SNATRuleSet `yaml:"rule_sets"`
}

// DestinationNAT is a set of NAT rules for DNAT.
type DestinationNAT struct {
	RuleSets []DNATRuleSet `yaml:"rule_sets"`
	Pools    []DNATPool    `yaml:"pools"`
}

type SNATRuleSet struct {
	Name  string     `yaml:"name"`
	From  string     `yaml:"from_zone"`
	To    string     `yaml:"to_zone"`
	Rules []SNATRule `yaml:"rules"`
}

type DNATRuleSet struct {
	Name  string     `yaml:"name"`
	From  string     `yaml:"from_zone"`
	To    string     `yaml:"to_zone"`
	Rules []DNATRule `yaml:"rules"`
}

// DNATRule is a single DNAT rule.
type DNATRule struct {
	Name  string        `yaml:"name"`
	Match DNATRuleMatch `yaml:"match"`
	Then  DNATRuleThen  `yaml:"then"`
}

// SecurityNATRule is a single SNAT rule.
type SNATRule struct {
	Name  string        `yaml:"name"`
	Match SNATRuleMatch `yaml:"match"`
}

// SecurityNATPool is a single NAT rule.
type DNATPool struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

// SNATRuleMatch is a match condition for a NAT rule.
type SNATRuleMatch struct {
	SourceAddressNames []string `yaml:"source_address_names"`
	SourceAddress      []string `yaml:"source_address"`
}

// DNATRuleMatch is a match condition for a NAT rule.
type DNATRuleMatch struct {
	DestinationAddressName string `yaml:"destination_address_name"`
	DestinationPort        string `yaml:"destination_port"`
	Protocol               string `yaml:"protocol"`
}

// DNATRuleMatch is a match condition for a NAT rule.
type DNATRuleThen struct {
	DestinationNAT struct {
		Pool string `yaml:"pool"`
	} `yaml:"destination_nat"`
}
