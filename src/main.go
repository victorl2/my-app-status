package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Service struct {
	Name   string
	Status string
}

var tmpl *template.Template

// initializes the template from static html files
func init() {
	tempTmpl, err := template.New("status.html").ParseFiles("templates/status.html")
	if err != nil {
		panic("could`t parse the template for the status page")
	}
	tmpl = tempTmpl
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	service := Service{Name: "Test", Status: "OK"}
	if err := tmpl.Execute(w, service); err != nil {
		json.NewEncoder(w).Encode(Service{Name: "Test", Status: err.Error()})
	}
}

//Starts a webserver to serve a simple html page
func main() {
	fmt.Println("################################")
	fmt.Println("#### Hello from MyAppServer ####")
	fmt.Println("################################")

	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe(":8080", nil)
}
