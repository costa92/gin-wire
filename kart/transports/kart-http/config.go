package kart_http

type HttpConfig struct {
	Name          string   `yaml:"name" json:"name" toml:"name"`
	Port          string   `yaml:"port" json:"port" toml:"port"`
	Mode          string   `yaml:"mode" json:"mode" toml:"mode"`
	ReadTimeout   int      `json:"read_timeout" yaml:"read_timeout" toml:"read_timeout"`
	WriteTimeout  int      `json:"write_timeout" yaml:"write_timeout" toml:"write_timeout"`
	Healthz       bool     `json:"healthz" yaml:"healthz" toml:"healthz"`
	EnableMetrics bool     `json:"enable-metrics" yaml:"enable-metrics" toml:"enable-metrics"`
	Middlewares   []string `json:"middlewares" yaml:"middlewares" toml:"middlewares"`
}
