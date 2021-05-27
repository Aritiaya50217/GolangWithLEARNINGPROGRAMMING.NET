package main

import (
	"fmt"
	"net/http"
	"test/apis/accountapi"
	"test/apis/demoapi"
	"test/middlewares/jwtauth"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/account/generatekey", accountapi.GenerateToken).Methods("POST")

	router.HandleFunc("/api/account/checktoken", accountapi.CheckToken).Methods("GET")

	// เรียกใช้ func JWTAuth โดยมีการ return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)
	// จึงสามารถเรียกใช้ func Demo1 ได้
	// Methods แบบ GET แสดงข้อความตามที่ระบุใน func Demo1
	router.Handle("/api/demo/demo1", jwtauth.JWTAuth(http.HandlerFunc(demoapi.Demo1))).Methods("GET")
	router.Handle("/api/demo/demo2", jwtauth.JWTAuth(http.HandlerFunc(demoapi.Demo2))).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
