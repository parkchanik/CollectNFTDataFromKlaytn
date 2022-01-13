package config

import (
	"encoding/json"
	"os"
)

type ConnectConfig struct {
	AccessKeyID     string `json:"AccessKeyID"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Endpoint        string `json:"Endpoint"`
	Network         string `json:"Network"`
}

func LoadConfigration(path string) (ConnectConfig, error) {
	var config ConnectConfig
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	return config, nil
}
