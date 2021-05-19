package home_controller

import (
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter,r*http.Request){
	tmp ,_ := template.ParseFiles("views/home/index.html")
	tmp.Execute(w,nil)
}