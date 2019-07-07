package main

import (
	"html/template"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("templates/crossoff.html"))
var validPath = regexp.MustCompile("^/(crossoff)")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", Item{Name: "bananas"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func crossoffHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "crossoff")
}