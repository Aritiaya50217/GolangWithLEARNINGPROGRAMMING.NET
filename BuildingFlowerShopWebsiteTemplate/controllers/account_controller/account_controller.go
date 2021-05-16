package account_controller

import (
	"html/template"
	"net/http"
)

func Myaccount(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/templates/mytemplate.html",
		"views/account/myaccount.html",
	)

	tmp.ExecuteTemplate(w, "layout", nil)
}

// สมัครสมาชิก
func Register(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/templates/mytemplate.html",
		"views/account/register.html",
	)
	tmp.ExecuteTemplate(w, "layout", nil)
}
