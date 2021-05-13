package main

import (
	// testwebsite คือ ชื่อmodlueที่เราสร้าง .mod
	"net/http"
	"testwebsite/controllers/demo_controller"
)

func main() {
	// การแสดงไฟล์ภาพ
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// controllers คือ ชื่อโฟล์เดอร์ที่เราสร้างในการเก็บ package
	// Index คือ function ในไฟล์ demo_controller
	http.HandleFunc("/", demo_controller.Index)
	http.HandleFunc("/demo/index", demo_controller.Index)

	// port
	http.ListenAndServe(":8080", nil)

}
