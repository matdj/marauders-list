package main

import (
	"fmt"
	"net/http"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "edit", Item{Name: "Cucumber"}, reloadTemplateFn())
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	fmt.Println(body)
	http.Redirect(w, r, "/crossoff", http.StatusFound)
}