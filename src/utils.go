package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config represents tool configurations
type Config struct {
	EncryptionKey string `json:"encryption_key"`
}

// LoadConfig reads configuration from a file
func LoadConfig(filename string) (*Config, error) {
	fmt.Println("Loading configuration file...")

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	fmt.Println("Configuration loaded successfully.")
	return &config, nil
}
