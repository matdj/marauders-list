package main

import (
	"marauders-list/controller"
	"marauders-list/htmlutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", controller.HealthHandler)
	http.HandleFunc("/crossoff", htmlutil.MakeHandler(controller.CrossoffHandler))
	http.HandleFunc("/edit", htmlutil.MakeHandler(controller.EditHandler))
	http.HandleFunc("/items", controller.SaveHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
