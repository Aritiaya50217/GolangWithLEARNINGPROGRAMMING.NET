package aboutus_controller

import (
	"html/template"
	"net/http"
)

func AboutUs(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		// หน้าหลัก
		"views/templates/mytemplate.html",
		"views/aboutus/aboutus.html",
	)
	tmp.ExecuteTemplate(w, "layout", nil)
}
