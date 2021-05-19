package category_controller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("views/admin/category/index.html")
	tmp.Execute(w, nil)
}
