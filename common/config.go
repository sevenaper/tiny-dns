package common

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

var DnsConfig *Config

// LoadConfig load the yaml config
func LoadConfig() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(data, DnsConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	return nil
}
