package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Uuid     string `json:"uuid"`
	CpuCount int    `json:"cpu_count"`
	Memory   int    `json:"memory"`
}

func SaveConfig(uuid string, cpus, mem int) error {
	path := os.ExpandEnv("$HOME/.dlite/config.json")
	output, err := os.Create(path)
	if err != nil {
		return err
	}

	defer output.Close()
	config := Config{
		Uuid:     uuid,
		CpuCount: cpus,
		Memory:   mem,
	}

	b, err := json.Marshal(config)
	if err != nil {
		return err
	}

	output.Write(b)
	return changePermissions(path)
}

func ReadConfig() (Config, error) {
	var config Config
	file, err := os.Open(os.ExpandEnv("$HOME/.dlite/config.json"))
	if err != nil {
		return Config{}, err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
