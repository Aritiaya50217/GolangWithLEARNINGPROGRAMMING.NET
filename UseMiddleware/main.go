package main

import (
	"net/http"
	"test/controllers/demo_controller"
	"test/controllers/middlewares"
)

func main() {

	http.HandleFunc("/", demo_controller.Index)
	http.HandleFunc("/demo", demo_controller.Index)
	http.HandleFunc("/demo/index", middlewares.Log1Middleware(demo_controller.Index))
	http.HandleFunc("/demo/index2", middlewares.Log2Middleware(demo_controller.Index2))
	http.HandleFunc("/demo/index3", middlewares.Log3Middleware(demo_controller.Index3))

	//port
	http.ListenAndServe(":8080", nil)
}
