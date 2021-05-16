package home_controller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles(
		// หน้าหลัก
		"views/templates/mytemplate.html",
		// ไฟล์ .html ที่ต้องการดึงข้อมูลมา
		"views/home/index.html",
	)
	tmp.ExecuteTemplate(w, "layout", nil)
}
