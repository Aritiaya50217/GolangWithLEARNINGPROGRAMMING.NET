package product_controller

import (
	"html/template"
	"net/http"
)

func Product(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("views/admin/product/product.html")
	tmp.Execute(w, nil)
}
