package main

import (
	"fmt"
)

func Demo1() {
	product := map[string]string{
		"id":          "p1",
		"name":        "name 1",
		"description": "Description 1",
	}
	fmt.Println(product)
	// เพิ่ม category ลงใน map ที่เราสร้าง
	product["category"] = "Category 1"
	// แสดงข้อมูลทั้งหมดใน product
	fmt.Println(product)
	fmt.Println("Product Info")
	// แสดง id
	fmt.Println("id:", product["id"])
	// แสดง name
	fmt.Println("name:", product["name"])
	// แสดง description
	fmt.Println("description:", product["description"])
	// แสดง category
	fmt.Println("category:", product["category"])
}

func Demo2() {
	// ฟังก์ชัน make built-in
	student := make(map[string]string)
	student["id"] = "st01"
	student["name"] = "name 1"
	fmt.Println(student)
}

func Demo3() {
	product := map[string]string{
		// key : value
		"id":          "p1",
		"name":        "name 1",
		"description": "Description 1",
	}
	// value เป็นตัวแปรเก็บค่า name 11
	// ok type bool
	// name 11 เป็น key ที่ใช้ในการตรวจสอบ

	value, ok := product["name 11"]
	if ok {
		// ถ้ามี name 11 ใน product
		// แสดง value
		fmt.Println(value)
	} else {
		// ถ้าไม่มี name 11 ใน product
		// แสดงคำว่า Invalid
		fmt.Println("Invalid")
	}

}

func Demo4() {
	//  สร้าง map เก็บค่าแบบ string ลงใน slice โดยใน slice ต้องเป็น string
	// slice เก็บข้อมูลกี่ชุดก็ได้ไม่จำกัด

	categories := map[string][]string{
		"category 1": {"product 1", "product 2"},
		"category 2": {"product 3", "product 4", "product 5"},
		"category 3": {"product 1"},
	}
	// แสดงข้อมูลทั้งหมดใน categories
	fmt.Println(categories)
	fmt.Println("Categories List")
	// แสดง key กับ value ตามลำดับที่อยู่ใน categories
	for key, value := range categories {
		// แสดง key
		fmt.Println(key)
		// สร้างตัวแปร v เก็บค่าลำดับของ value
		for _, v := range value {
			// "/t" คือ การ tab
			fmt.Println("\t", v)
		}
	}
}

func Demo5() {
	// interface ใน Go เป็น m:n
	product := map[string]interface{}{
		"id":       "p1",
		"name":     "name 1",
		"price":    4.5,
		"quantity": 6,
		"status":   true,
	}
	fmt.Println(product)
	fmt.Println("Product Info")
	for key, value := range product {
		fmt.Println(key, ":", value)
	}
}

func main() {
	// Demo1()
	// Demo2()
	// Demo3()
	// Demo4()
	Demo5()
}
