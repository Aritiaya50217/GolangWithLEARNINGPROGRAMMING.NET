package main

import (
	"net/http"
	"test/controllers/demo_controller"
)

func main() {
	http.HandleFunc("/", demo_controller.Index)
	http.HandleFunc("/demo", demo_controller.Index)

	//port
	http.ListenAndServe(":8080", nil)
}
