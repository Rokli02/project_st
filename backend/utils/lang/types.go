package lang

type Language struct {
	User    TextMap `yaml:"USER"`
	Common  TextMap `yaml:"COMMON"`
	Unknown string  `yaml:"UNKNOWN"`
}

type TextMap map[string]string
