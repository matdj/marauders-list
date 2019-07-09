package controller

import (
	"fmt"
	"marauders-list/domain"
	"marauders-list/htmlutil"
	"net/http"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	htmlutil.RenderTemplate(w, "edit", domain.Item{Name: "Cucumber"}, htmlutil.ReloadTemplateFn())
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	fmt.Println(body)
	http.Redirect(w, r, "/crossoff", http.StatusFound)
}
