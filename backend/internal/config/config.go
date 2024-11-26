package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct of app's config
type Config struct {
	App struct {
		Mock      bool   `yaml:"mock"`
		LogLevel  string `yaml:"log_level"`
		LogFormat string `yaml:"log_format"`

		HTTP struct {
			ListenAddr string `yaml:"listen_addr"`
		} `yaml:"http"`
	} `yaml:"app"`
}

// MustLoad load config from file or panic if error
func MustLoad(configPath string) *Config {
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return &config
}
