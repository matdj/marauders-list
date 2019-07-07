package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var supportedPaths = regexp.MustCompile("^/(crossoff|edit)")

func reloadTemplateFn() func() *template.Template {
	return func() *template.Template { return template.Must(template.ParseFiles("templates/crossoff.html", "templates/edit.html")) }
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/crossoff", makeHandler(crossoffHandler))
	http.HandleFunc("/edit", makeHandler(editHandler))
	http.HandleFunc("/items", saveHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
