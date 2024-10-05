package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port     string `json:"port"`
	DataPath string `json:"data_path"`
	JWTSecret string `json:"jwt_secret"`
}

func LoadConfig() *Config {
	file, err := os.Open("src/config/config.json")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	return &cfg
}
