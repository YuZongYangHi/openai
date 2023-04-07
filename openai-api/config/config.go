package config

import "github.com/YuZongYangHi/chatgpt-proxy/openai-api/pkg/util/parsers"

var config *Config

type Config struct {
	Proxy ProxyConfig `yaml:"proxy"`
	DB    DBConfig    `yaml:"db"`
	HTTP  HTTPConfig  `yaml:"http"`
}

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
}

type ProxyConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

type DBConfig struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	Password string `yaml:"password"`
}

func AppConfig() *Config {
	return config
}

func NewConfig(in string) error {
	return parsers.ParserConfigurationByFile(parsers.YAML, in, &config)
}
