package main

import (
	"net/http"
	"test/controllers/product_controller"
)

func main() {

	http.HandleFunc("/", product_controller.Index)
	http.HandleFunc("/product", product_controller.Index)
	http.HandleFunc("/product/index", product_controller.Index)
	http.HandleFunc("/product/add", product_controller.Add)
	// เพิ่มข้อมูลใน database
	http.HandleFunc("/product/processadd", product_controller.ProcessAdd)
	// ลบข้อมูลใน database
	http.HandleFunc("/product/delete", product_controller.Delete)
	// แก้ไขข้อมูลใน database
	http.HandleFunc("product/edit",product_controller.Edit)
	
	
	// port
	http.ListenAndServe(":8080", nil)
}
