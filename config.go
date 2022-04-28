package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	LOG_PATHS []string
}

func loadConfig() Config {
	var config Config
	configFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer configFile.Close()

	jsonparser := json.NewDecoder(configFile)
	jsonparser.Decode(&config)

	return config
}
