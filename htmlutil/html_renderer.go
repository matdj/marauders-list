package htmlutil

import (
	"html/template"
	"net/http"
	"regexp"
)

var supportedPaths = regexp.MustCompile("^/(crossoff|edit)")

func ReloadTemplateFn() func() *template.Template {
	return func() *template.Template {
		return template.Must(template.ParseFiles("templates/crossoff.html", "templates/edit.html"))
	}
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := supportedPaths.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}, reloadTemplateFn func() *template.Template) {
	//TODO move to global to cache templates
	var templates = reloadTemplateFn()

	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
