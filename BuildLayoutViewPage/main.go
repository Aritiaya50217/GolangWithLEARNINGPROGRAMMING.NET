package main

import (
	"net/http"
	"test/controllers/home_controller"
	"test/controllers/aboutus_controller"
	"test/controllers/news_controller"
)
func main() {
	http.HandleFunc("/",home_controller.Index)
	http.HandleFunc("/home",home_controller.Index)
	//http.HandleFunc("/home/index",home_controller.Index)

	http.HandleFunc("/aboutus",aboutus_controller.AboutUs)
	//http.HandleFunc("/aboutus/aboutus",aboutus_controller.AboutUs)

	http.HandleFunc("/news",news_controller.News)
	//http.HandleFunc("/news/news",news_controller.News)



	// port
	http.ListenAndServe(":8080",nil)
}