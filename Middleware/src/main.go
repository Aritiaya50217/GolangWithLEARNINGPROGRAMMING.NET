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
	router.Handle("/api/demo1", middlewares.Middleware1(http.HandlerFunc(apis.Demo1API))).Methods("GET")

	router.Handle("/api/demo2", middlewares.Middleware2(http.HandlerFunc(apis.Demo2API))).Methods("GET")

	router.Handle("/api/demo3", middlewares.Middleware3_1(middlewares.Middleware3_2(http.HandlerFunc(apis.Demo3API)))).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
