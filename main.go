package main

import (
	"fmt"
	app "myappstatus/app"
)

//Starts a webserver to serve a simple html page
func main() {
	fmt.Println("################################")
	fmt.Println("#### Hello from MyAppStatus ####")
	fmt.Println("################################")

	app.StartServer()
}
