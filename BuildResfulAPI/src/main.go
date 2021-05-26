package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test/entities"

	"github.com/gorilla/mux"
)

func Demo1API(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Demo 1")
}
func Demo2API(w http.ResponseWriter, r *http.Request) {
	// set ให้แสดงผลแบบ html
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<b><u><i>Demo 2</i></u></b>")
}

func HelloAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fullName := vars["fullName"]

	fmt.Fprint(w, "Hello ", fullName)
}

func SumAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]
	// ตัวแปร number1 เก็บข้อมูลแบบ Int
	number1, err1 := strconv.ParseInt(a, 10, 64)
	number2, err2 := strconv.ParseInt(b, 10, 64)
	if err1 != nil || err2 != nil {
		fmt.Fprint(w, "Error 1:", err1, "Error 2:", err2)
	} else {
		fmt.Fprint(w, number1+number2)
	}
}

// การเพิ่มข้อมูล
func Create(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Println("Product Info")
		fmt.Println("Id: ", product.Id)
		fmt.Println("Name: ", product.Name)
		fmt.Println("Price: ", product.Price)
		fmt.Println("Quantity: ", product.Quantity)
		fmt.Println("Status: ", product.Status)
	}
}

// การแก้ไขข้อมูล
func Update(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Println("Product Info Update")
		fmt.Println("Id: ", product.Id)
		fmt.Println("Name: ", product.Name)
		fmt.Println("Price: ", product.Price)
		fmt.Println("Quantity: ", product.Quantity)
		fmt.Println("Status: ", product.Status)
	}
}

// การลบข้อมูลตาม id
func Delete(w http.ResponseWriter,r*http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w,"Id: ",id)
}

func main() {
	// conneted router
	r := mux.NewRouter()
	r.HandleFunc("/api/demo/demo1", Demo1API).Methods("GET")
	r.HandleFunc("/api/demo/demo2", Demo2API).Methods("GET")

	r.HandleFunc("/api/demo/hello/{fullName}", HelloAPI).Methods("GET")

	r.HandleFunc("/api/demo/sum/{a}/{b}", SumAPI).Methods("GET")
	// POST
	r.HandleFunc("/api/demo/create", Create).Methods("POST")
	// PUT
	r.HandleFunc("/api/demo/update", Update).Methods("PUT")
	// delete
	r.HandleFunc("/api/demo/delete/{id}",Delete).Methods("DELETE")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}

}
