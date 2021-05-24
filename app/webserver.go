package app

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpl *template.Template
var configuration *Config

//initializes the template from static html files
func initTemplate() *template.Template {
	tempTmpl, err := template.New("status.html").ParseFiles("templates/status.html")
	if err != nil {
		log.Println("couldn`t parse the template for the status page." + err.Error())
		return nil
	}
	return tempTmpl
}

//request handler for serving the status homepage
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, *configuration); err != nil {
		log.Println(err.Error())
		json.NewEncoder(w).Encode(configuration)
	}
}

func statusUpdater(waitInterval int) {
	go func() {
		for {
			ValidateServices(configuration)
			time.Sleep(time.Duration(waitInterval*60) * time.Second)
		}
	}()
}

//StartServer initializes the webserver serving a simple html page
func StartServer() {
	configuration = LoadGlobalConfig()
	tmpl = initTemplate()

	if configuration == nil || tmpl == nil {
		panic("The initial setup failed, check if there is a problem with the app configuration or template values")
	}

	statusUpdater(configuration.IntervalCheck)

	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe(":8080", nil)
}
