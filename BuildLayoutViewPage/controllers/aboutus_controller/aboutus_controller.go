package aboutus_controller

import (
	"net/http"
	"html/template"
)
var templates *template.Template
func init(){
	templates = template.Must(template.ParseFiles(
		// หน้าที่แสดง template หน้าหลัก
		"views/templates/mytemplate.html",
		// แสดงตามข้อมูลที่เราสร้างไว้ใน aboutus.html
		"views/aboutus_controller/aboutus.html",
	))
}
func AboutUs(w http.ResponseWriter,r*http.Request){
	// สร้าง map เก็บข้อความว่า Content of About Us Page 
	// เวลานำไปใช้ให้นำ content ไปใส่ใน .html
	data := map[string]interface{}{
		"content" : "Content of About Us Page",
	}
	// layout คือ คำที่ใช้ดึงข้อมูลมาแสดงในหน้าหลัก 
	// โดยจะใช้ แท็กว่า {{define "layout" }}
	templates.ExecuteTemplate(w, "layout", data)
}