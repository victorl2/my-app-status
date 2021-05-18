package app

import (
	"log"
	"net/http"
)

//SimpleHealthCheck sends http get request to the provided url
//if the status code is in the 200 range a true bool value is retuned indicating sucess, if not a false value is returned
func SimpleHealthCheck(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		log.Print("Request failed:" + err.Error())
		return false
	}
	return response.StatusCode >= 200 && response.StatusCode < 300
}
