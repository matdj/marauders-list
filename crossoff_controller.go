package main

import (
	"net/http"
)

var exampleShoppingList = ShoppingList{Items: []Item{Item{Name: "bananas"}}, CrossedOffItems: []Item{Item{Name: "strawberries"}}}

func crossoffHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "crossoff", exampleShoppingList, reloadTemplateFn())
}
