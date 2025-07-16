package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

const configPath = "config.json"

var config Config
var mutex sync.RWMutex

type Config struct {
	Port     string `json:"port" env:"PORT"`
	HTMLDir  string `json:"html_dir" env:"HTML_DIR"`
	Template string `json:"template" env:"TEMPLATE"`
}

func LoadConfig() error {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}

	port := os.Getenv("PORT")
	if port != "" {
		config.Port = port
	} else {
		config.Port = "8080"
	}

	htmlDir := os.Getenv("HTML_DIR")
	if htmlDir != "" {
		config.HTMLDir = htmlDir
	} else {
		config.HTMLDir = "html"
	}

	template := os.Getenv("TEMPLATE")
	if template != "" {
		config.Template = template
	}

	return nil
}

func GetConfig() *Config {
	mutex.RLock()
	defer mutex.RUnlock()

	return &config
}
