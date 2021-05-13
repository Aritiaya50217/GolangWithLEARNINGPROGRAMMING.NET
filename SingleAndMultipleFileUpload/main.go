package main

import (
	"net/http"
	"website/controllers/demo_controller"
)

func main() {
	http.HandleFunc("/demo/index", demo_controller.Index)
	// upload single file
	http.HandleFunc("/demo/singleupload", demo_controller.SingleUpload)
	http.HandleFunc("/demo/index2", demo_controller.Index2)
	// upload multiple file
	http.HandleFunc("/demo/multipleupload", demo_controller.MultipleUpload)
	// port
	http.ListenAndServe(":8080", nil)
}
