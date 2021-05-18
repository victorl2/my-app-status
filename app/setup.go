package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	IntervalCheck int       `json:"interval_check"`
	Subtitle      string    `json:"subtitle"`
	AdvancedCheck bool      `json:"advanced_check"`
	EnvsNames     []string  `json:"environments"`
	Apps          []Service `json:"services"`
}

type Service struct {
	Name string        `json:"name"`
	Envs []Environment `json:"environments"`
}

type Environment struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status int
}

func LoadGlobalConfig() *Config {
	var config Config
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Println("Failed to load configuration, " + err.Error())
		return nil
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	fmt.Printf("%+v\n", config)
	return &config
}

func validateServices(config *Config) bool {
	if config == nil {
		return false
	}
}
