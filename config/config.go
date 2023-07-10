package config

type Configuration struct {
	Server ServerConfig `yaml:"server" json:"server" toml:"server"`
}

type ServerConfig struct {
	Name string `yaml:"name" json:"name" toml:"name"`
	Addr string `yaml:"addr" json:"addr" toml:"addr"`
	Port string `yaml:"port" json:"port" toml:"port"`
}
