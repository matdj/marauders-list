package main

import (
	"log"
	"marauders-list/controller"
	"marauders-list/htmlutil"
	"marauders-list/services"
	"net/http"
	"os"
)

func determineRunMode() string {
	if os.Getenv("RUN_MODE") == "prod" {
		return ":8080"
	}
	return "localhost:8080"
}

func main() {
	service := new(services.ShoppingListService)
	edit := controller.EditController{Service: service}
	crossoff := controller.CrossoffController{Service: service}

	http.HandleFunc("/health", controller.HealthHandler)
	http.HandleFunc("/crossoff", htmlutil.MakeHandler(crossoff.CrossoffHandler))
	http.HandleFunc("/edit", htmlutil.MakeHandler(edit.EditHandler))
	http.HandleFunc("/add", htmlutil.MakeHandler(edit.EditHandler))
	http.HandleFunc("/items", edit.SaveHandler)

	runMode := determineRunMode()

	log.Println("Starting webserver " + runMode)
	log.Fatal(http.ListenAndServe(runMode, nil))
}
