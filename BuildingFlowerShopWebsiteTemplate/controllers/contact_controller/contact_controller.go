package contact_controller

import (
	"html/template"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/templates/mytemplate.html",
		"views/contact/contact.html",
	)
	tmp.ExecuteTemplate(w, "layout", nil)
}
