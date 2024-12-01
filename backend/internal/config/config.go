package config

import (
	"fmt"
	"os"
	"time"

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

	HTTPServer HTTPServer `yaml:"http_server"`

	HTTPClient HTTPClient `yaml:"http_client"`
}

type HTTPServer struct {
	ListenAddr   string        `yaml:"listen_addr"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type HTTPClient struct {
	Timeout time.Duration `yaml:"timeout"`
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

	fmt.Printf("config: %+v\n", config)

	return &config
}
