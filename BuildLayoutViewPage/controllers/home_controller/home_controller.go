package home_controller

import (
	"html/template"
	"net/http"
)

// สร้างตัวแปร เก็บค่า Package template
var templates *template.Template

func init() {
	// การเปิดไฟล์ที่ใช้เป็น template ในแต่ละโฟลเดอร์
	templates = template.Must(template.ParseFiles(
		// template หน้าหลัก
		"views/templates/mytemplate.html",
		// แสดงตามข้อมูลที่เราสร้างไว้ใน index.html
		"views/home_controller/index.html",
	))
}
func Index(w http.ResponseWriter, r *http.Request) {
	// layout คือ คำที่ใช้ดึงข้อมูลมาแสดงในหน้าหลัก 
	// โดยจะใช้ แท็กว่า {{define "layout" }}
	templates.ExecuteTemplate(w, "layout", nil)
}
