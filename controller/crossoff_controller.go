package controller

import (
	"marauders-list/domain"
	"marauders-list/htmlutil"
	"marauders-list/services"
	"net/http"
)

var service = new(services.ShoppingListService)

func CrossoffHandler(w http.ResponseWriter, r *http.Request) {
	addWeetbix()

	var exampleShoppingList = domain.ShoppingList{
		Items:           service.Items(),
		CrossedoffItems: service.CrossedoffItems()}

	htmlutil.RenderTemplate(w, "crossoff", exampleShoppingList, htmlutil.ReloadTemplateFn())
}

func addWeetbix() {
	weetbix := domain.Item{Name: "weetbix"}
	if !service.ItemsContains(weetbix) {
		service.Add(weetbix)
	}
}
