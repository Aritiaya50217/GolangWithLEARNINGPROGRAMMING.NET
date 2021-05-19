package demo_controller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/demo/index.html",
	)
	tmp.Execute(w, nil)
}
func Index2(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/demo/index2.html",
	)
	tmp.Execute(w, nil)
}
func Index3(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/demo/index3.html",
	)
	tmp.Execute(w, nil)
}
