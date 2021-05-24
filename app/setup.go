package app

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Config struct {
	IntervalCheck int       `json:"interval_check"`
	Subtitle      string    `json:"subtitle"`
	AdvancedCheck bool      `json:"advanced_check"`
	EnvsNames     []string  `json:"environments"`
	Apps          []Service `json:"services"`
	LogoURL       string    `json:"logo_url"`     //redirect url when click on the logo
	LogoSRCURL    string    `json:"logo_src_url"` //url to the logo
	GlobalStatus  int
}

func (config *Config) UpdateAppStatus(serviceName string, envIndex, newStatus int) {
	for _, service := range config.Apps {
		if service.Name == serviceName && service.Envs[envIndex].Status != newStatus {
			service.Envs[envIndex].Status = newStatus
			log.Printf("%v(%v) status updated to: %v\n", service.Name, config.EnvsNames[envIndex], service.Envs[envIndex].Status)
		}
	}
}

func (config *Config) UpdateGlobalStatus(newStatus int) {
	if config.GlobalStatus != newStatus {
		config.GlobalStatus = newStatus
		log.Printf("Global application status changed to %v", newStatus)
	}
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

	loadedConfig := &config
	if !validateConfig(loadedConfig) {
		return nil
	}
	return loadedConfig
}

func validateConfig(config *Config) bool {
	if config == nil {
		log.Print("Invalid configuration, empty information provided")
		return false
	}

	for _, service := range config.Apps {
		for _, environment := range service.Envs {
			if !strings.Contains(environment.URL, "http://") && !strings.Contains(environment.URL, "https://") {
				log.Printf("The service %s has a invalid URL in the configuration file, it dont contain the prefix http or https", service.Name)
				return false
			}
		}
	}
	return true
}
