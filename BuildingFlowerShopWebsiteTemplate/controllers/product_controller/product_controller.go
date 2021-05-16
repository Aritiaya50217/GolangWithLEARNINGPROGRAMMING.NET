package product_controller

import (
	"html/template"
	"net/http"
)

func Category(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/templates/mytemplate.html",
		"views/product/category.html",
	)
	tmp.ExecuteTemplate(w, "layout", nil)
}
func Specials(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/templates/mytemplate.html",
		"views/product/specials.html",
	)
	tmp.ExecuteTemplate(w, "layout", nil)
}

func Details(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/templates/mytemplate.html",
		"views/product/details.html",
	)
	tmp.ExecuteTemplate(w, "layout", nil)
}
