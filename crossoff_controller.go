package main

import (
	"html/template"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("templates/crossoff.html"))
var validPath = regexp.MustCompile("^/(crossoff)")

var exampleShoppingList = ShoppingList{Items: []Item{Item{Name: "bananas"}}, CrossedOffItems: []Item{Item{Name: "strawberries"}}}

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
	err := templates.ExecuteTemplate(w, tmpl+".html", exampleShoppingList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func crossoffHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "crossoff")
}