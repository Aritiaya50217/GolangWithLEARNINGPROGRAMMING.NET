package main

import (
	"fmt"
	"net/http"
	"test/apis"
	"test/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// BasicAuth(apis.FindAll) เรียกใช้ func FindAll ที่อยู่ในโฟล์เดอร์ apis
	router.Handle("/api/product/findall", middlewares.BasicAuth(apis.FindAll)).Methods("GET")

	router.Handle("/api/product/search", middlewares.BasicAuth(apis.Search)).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
