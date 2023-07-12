package kart_http

type HttpConfig struct {
	Name         string `yaml:"name" json:"name" toml:"name"`
	Host         string `yaml:"host" json:"host" toml:"host"`
	Port         string `yaml:"port" json:"port" toml:"port"`
	Mode         string `yaml:"mode" json:"mode" toml:"mode"`
	ReadTimeout  int    `json:"read_timeout" yaml:"read_timeout" toml:"read_timeout" json:"readTimeout,omitempty"`
	WriteTimeout int    `json:"write_timeout" yaml:"write_timeout" toml:"write_timeout" json:"writeTimeout,omitempty"`
}
