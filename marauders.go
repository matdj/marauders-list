package main

import (
	"log"
	"marauders-list/controller"
	"marauders-list/htmlutil"
	"marauders-list/services"
	"net/http"
)

func main() {
	service := new(services.ShoppingListService)
	edit := controller.EditController{Service: service}
	crossoff := controller.CrossoffController{Service: service}

	http.HandleFunc("/health", controller.HealthHandler)
	http.HandleFunc("/crossoff", htmlutil.MakeHandler(crossoff.CrossoffHandler))
	http.HandleFunc("/edit", htmlutil.MakeHandler(edit.EditHandler))
	http.HandleFunc("/items", edit.SaveHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
