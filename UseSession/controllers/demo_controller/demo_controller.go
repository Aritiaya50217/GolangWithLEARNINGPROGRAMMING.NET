package demo_controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"test/controllers/entities"

	"github.com/gorilla/sessions"
)

// สร้างตัวแปร เก็บค่า sessions
var store = sessions.NewCookieStore([]byte("mykey"))

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session1")
	session.Values["age"] = 23
	session.Values["username"] = "abc"

	// นำ Product struct ที่เราสร้างในโฟลเดอ์ entities มาแสดง
	product := entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    1.99,
		Quantity: 5,
		Status:   true,
	}
	strProduct, _ := json.Marshal(product)
	//
	session.Values["product"] = string(strProduct)

	// แสดงข้อมูลหลายๆชุด
	// เก็บข้อมูลใส่ใน slice
	products := []entities.Product{
		entities.Product{
			Id:       "p01",
			Name:     "name 1",
			Price:    1.99,
			Quantity: 5,
			Status:   true,
		},
		entities.Product{
			Id:       "p02",
			Name:     "name 2",
			Price:    2.99,
			Quantity: 2,
			Status:   true,
		},
		entities.Product{
			Id:       "p03",
			Name:     "name 3",
			Price:    3.99,
			Quantity: 3,
			Status:   true,
		},
	}
	strProducts, _ := json.Marshal(products)
	session.Values["products"] = string(strProducts)

	// save และแสดงผล
	session.Save(r, w)

	tmp, _ := template.ParseFiles(
		"views/demo/index.html",
	)
	tmp.Execute(w, nil)
}

func Display1(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session1")
	// เก็บ session มาจากค่า username โดย ค่า username เก็บข้อมูลแบบ string
	username := session.Values["username"]
	age := session.Values["age"]

	fmt.Println("username: ", username)
	fmt.Println("age: ", age)

	// สร้าง map เก็บข้อมูล
	// เพื่อนำไปใช้ใน .html
	data := map[string]interface{}{
		"username": username,
		"age":      age,
	}
	tmp, _ := template.ParseFiles(
		"views/demo/display1.html",
	)
	tmp.Execute(w, data)
}
func Display2(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session1")
	// เก็บค่าเป็น string
	strProduct := session.Values["product"].(string)
	// สร้างตัวแปร เก็บข้อมูลของ Product struct ในโฟลเดอร์ entities
	var product entities.Product

	json.Unmarshal([]byte(strProduct), &product)

	fmt.Println("Producy Info")
	fmt.Println("id : ", product.Id)
	fmt.Println("name : ", product.Name)
	fmt.Println("price : ", product.Price)
	fmt.Println("quantity : ", product.Quantity)
	fmt.Println("status : ", product.Status)

	// สร้าง map เก็บข้อมูล
	// เพื่อนำไปใช้ใน .html
	data := map[string]interface{}{
		"product": product,
	}
	tmp, _ := template.ParseFiles("views/demo/display2.html")
	tmp.Execute(w, data)
}
func Display3(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session1")

	strProducts := session.Values["products"].(string)
	// เก็บค่าใน slice สำหรับแสดงข้อมูลหลายชุด
	var products []entities.Product
	json.Unmarshal([]byte(strProducts), &products)

	fmt.Println("Product List")
	// แสดงข้อมูลทั้งหมด ตามที่ระบุใน Product struct ออกมาทีละชุดตามลำดับ
	for _, product := range products {
		fmt.Println("id : ", product.Id)
		fmt.Println("name : ", product.Name)
		fmt.Println("price : ", product.Price)
		fmt.Println("quantity : ", product.Quantity)
		fmt.Println("status : ", product.Status)
		fmt.Println("*************************")
	}
	data := map[string]interface{}{
		"products": products,
	}
	tmp, _ := template.ParseFiles(
		"views/demo/display3.html",
	)
	tmp.Execute(w, data)
}
