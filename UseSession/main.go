package main

import (
	"net/http"
	"test/controllers/demo_controller"
)

func main() {

	http.HandleFunc("/", demo_controller.Index)
	http.HandleFunc("/demo", demo_controller.Index)
	http.HandleFunc("/demo/display1", demo_controller.Display1)
	http.HandleFunc("/demo/display2", demo_controller.Display2)
	http.HandleFunc("/demo/display3", demo_controller.Display3)

	// port
	http.ListenAndServe(":8080", nil)
}
