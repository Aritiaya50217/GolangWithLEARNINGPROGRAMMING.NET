package main

import (
	"net/http"
	"test/controllers/aboutus_controller"

	"test/controllers/account_controller"
	"test/controllers/contact_controller"
	"test/controllers/home_controller"
	"test/controllers/product_controller"
)

func main() {
	// เปิดไฟล์ static
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", home_controller.Index)
	http.HandleFunc("/home", home_controller.Index)
	http.HandleFunc("/home/index", home_controller.Index)

	http.HandleFunc("/aboutus", aboutus_controller.AboutUs)
	http.HandleFunc("/aboutus/aboutus", aboutus_controller.AboutUs)

	http.HandleFunc("/product/category", product_controller.Category)
	http.HandleFunc("/product/specials", product_controller.Specials)
	http.HandleFunc("/product/details", product_controller.Details)

	http.HandleFunc("/account/myaccount", account_controller.Myaccount)
	http.HandleFunc("/account/register", account_controller.Register)

	http.HandleFunc("/contact", contact_controller.Contact)

	// port
	http.ListenAndServe(":8080", nil)

}
