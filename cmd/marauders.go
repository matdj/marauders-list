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

func tryDefaultEnvVar(envVar string, envVarDefault string) string {
	val := envVarDefault
	if os.Getenv(envVar) != "" {
		val = os.Getenv(envVar)
	}
	return val
}

func main() {
	templatesDir := tryDefaultEnvVar("TEMPLATES_DIR", "templates")

	htmlRenderer := htmlutil.HtmlRenderer{TemplatesDir: templatesDir}
	service := new(services.ShoppingListService)
	edit := controller.EditController{Service: service, HtmlRenderer: &htmlRenderer}
	crossoff := controller.CrossoffController{Service: service, HtmlRenderer: &htmlRenderer}
	health := controller.HealthController {}

	http.HandleFunc("/health", health.HealthHandler)
	http.HandleFunc("/crossoff", htmlutil.MakeHandler(crossoff.CrossoffHandler))
	http.HandleFunc("/edit", htmlutil.MakeHandler(edit.EditHandler))
	http.HandleFunc("/add", htmlutil.MakeHandler(edit.EditHandler))
	http.HandleFunc("/items", edit.SaveHandler)

	runMode := determineRunMode()

	log.Println("Starting webserver " + runMode)
	log.Fatal(http.ListenAndServe(runMode, nil))
}
