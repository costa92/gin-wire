package config

import (
	"costa92/gin-wire/pkg"
)

type Configuration struct {
	Server   ServerConfig    `yaml:"server" json:"server" toml:"server"`
	MasterDB pkg.MySQLConfig `yaml:"master_db" json:"master_db" toml:"master_db"`
}

type ServerConfig struct {
	Name string `yaml:"name" json:"name" toml:"name"`
	Addr string `yaml:"addr" json:"addr" toml:"addr"`
	Port string `yaml:"port" json:"port" toml:"port"`
}
