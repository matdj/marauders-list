package controller

import (
	"marauders-list/domain"
	"marauders-list/htmlutil"
	"marauders-list/services"
	"net/http"
)

type CrossoffController struct {
	Service      *services.ShoppingListService
	HtmlRenderer *htmlutil.HtmlRenderer
}

func (c *CrossoffController) CrossoffHandler(w http.ResponseWriter, r *http.Request) {
	c.addWeetbix()

	var exampleShoppingList = domain.ShoppingList{
		Items:           c.Service.Items(),
		CrossedoffItems: c.Service.CrossedoffItems()}

	htmlutil.RenderTemplate(w, "crossoff", exampleShoppingList, c.HtmlRenderer.ReloadTemplateFn())
}

func (c *CrossoffController) addWeetbix() {
	weetbix := domain.Item{Name: "weetbix"}
	if !c.Service.ItemsContains(weetbix) {
		c.Service.Add(weetbix)
	}
}
