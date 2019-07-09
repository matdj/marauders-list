package controller

import (
	"marauders-list/domain"
	"marauders-list/htmlutil"
	"marauders-list/services"
	"net/http"
)

type EditController struct {
	Service *services.ShoppingListService
}

func (controller *EditController) EditHandler(w http.ResponseWriter, r *http.Request) {
	htmlutil.RenderTemplate(w, "edit", domain.Item{Name: "Cucumber"}, htmlutil.ReloadTemplateFn())
}

func (controller *EditController) SaveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	controller.Service.Add(domain.Item{Name: body})
	http.Redirect(w, r, "/crossoff", http.StatusFound)
}
