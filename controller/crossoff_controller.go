package controller

import (
	"marauders-list/htmlutil"
	"marauders-list/domain"
	"net/http"
)

var exampleShoppingList = domain.ShoppingList{Items: []domain.Item{domain.Item{Name: "bananas"}}, 
	CrossedOffItems: []domain.Item{domain.Item{Name: "strawberries"}}}

func CrossoffHandler(w http.ResponseWriter, r *http.Request) {
	htmlutil.RenderTemplate(w, "crossoff", exampleShoppingList, htmlutil.ReloadTemplateFn())
}
