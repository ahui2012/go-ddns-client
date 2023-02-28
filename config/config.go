package config

import (
	"encoding/json"
	"os"
)

type DomainConfig struct {
	Domain     string
	Provider   string
	SecretId   string
	SecretKey  string
	SubDomains []string
}

type AppConfig struct {
	UpdateInterval uint32
	PubIPUrls      []string
	Domains        []DomainConfig
}

func GetConfig(filePath string) (*AppConfig, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config AppConfig
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
