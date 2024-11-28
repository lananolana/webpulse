package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct of app's config
type Config struct {
	App App `yaml:"app"`
}

type App struct {
	Mock      bool   `yaml:"mock"`
	LogLevel  string `yaml:"log_level"`
	LogFormat string `yaml:"log_format"`

	Test string `yaml:"test"`

	HTTP HTTP `yaml:"http"`
}

type HTTP struct {
	ListenAddr string `yaml:"listen_addr"`
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

	fmt.Println(config.App.Test)

	fmt.Printf("config: %+v\n", config)

	return &config
}
