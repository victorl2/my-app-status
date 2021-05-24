package app

import (
	"log"
	"net/http"
	"time"
)

func ValidateServices(config *Config) {
	log.Println("Verifying status of services")
	var status = 0
	for _, service := range config.Apps {
		for i, environment := range service.Envs {
			var newStatus int
			if !config.AdvancedCheck {
				newStatus = simpleHealthCheck(environment.URL)
			} else {
				newStatus = advancedHealthCheck(environment.URL)
			}
			config.UpdateAppStatus(service.Name, i, newStatus)
			status += newStatus
		}
	}

	if status > 1 {
		config.UpdateGlobalStatus(2)
	} else if status == 1 {
		config.UpdateGlobalStatus(1)
	} else {
		config.UpdateGlobalStatus(0)
	}
}

//SimpleHealthCheck sends http get request to the provided url
//if the status code is in the 200 range a true bool value is retuned indicating sucess, if not a false value is returned
func simpleHealthCheck(url string) int {
	start := time.Now()
	response, err := http.Get(url)

	//Degraded
	if time.Since(start) > (time.Second * 10) {
		return 1
	}

	//Offline
	if err != nil || response.StatusCode < 200 || response.StatusCode >= 300 {
		if err != nil {
			log.Print("Request failed:" + err.Error())
		}
		return 2
	}

	return 0
}

func advancedHealthCheck(url string) int {
	return 2
}
