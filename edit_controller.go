package main

import (
	"net/http"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "edit", Item{Name: "Cucumber"}, reloadTemplateFn())
}

