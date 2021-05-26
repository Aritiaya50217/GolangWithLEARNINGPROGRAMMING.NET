package main

import (
	"encoding/json"
	"fmt"
	"test/entities"
)

func Demo1() {
	product := entities.Product{"p1", "name 1", 2, 5, true, entities.Category{"c1", "Category 1"}, []entities.Comment{
		entities.Comment{"title 1", "content 1"},
		entities.Comment{"title 2", "content 2"},
		entities.Comment{"title 3", "content 3"},
	}}
	result, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}

func Demo2() {
	var str string = `{
		"id":"p1",
		"name":"name 1 ", 
		"price" :2,
		"quantity":5, 
		"status":true
	}`
	var product entities.Product
	json.Unmarshal([]byte(str), &product)
	fmt.Println("Product Info")
	fmt.Println(product.ToString())
}

func Demo3() {
	product := entities.Product{
		Id:       "p1",
		Name:     "name 1",
		Price:    2,
		Quantity: 5,
		Status:   true,
		Category: entities.Category{
			Id:   "c1",
			Name: "Category 1",
		},
	}
	result, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}

func Demo4() {
	var str string = `{
		"id" : "p1",
		"name"	: "name 1",
		"price":4.5,
		"quantity":8,
		"status" : "true",
		"Category":{
			"id" :"c1",
			"name" :"Category 1"
		}}`
	var product entities.Product
	json.Unmarshal([]byte(str), &product)
	fmt.Println("Product Info")
	fmt.Println(product.ToString())
	fmt.Println("Category Info")
	fmt.Println("\tCategory Id :", product.Category.Id)
	fmt.Println("\tCategory Id :", product.Category.Name)
}

func Demo5() {
	product := entities.Product{
		Id:       "p1",
		Name:     "name 1",
		Price:    2,
		Quantity: 5,
		Status:   true,
		Category: entities.Category{
			Id:   "c1",
			Name: "Category 1",
		},
		Comments: []entities.Comment{
			entities.Comment{Title: "title 1", Content: "content 1"},
			entities.Comment{Title: "title 2", Content: "content 2"},
			entities.Comment{Title: "title 3", Content: "content 3"},
		},
	}
	result, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}

func Demo6() {
	var str string = `{
		"id" :"p1",
		"name":"name 1",
		"price" : 4.5,
		"quantity" : 8,
		"status" : "true",
		"category" :{
			"id":"c1",
			"name":"Category 1"
		},
		"comments" : [
			{"title":"title 1","content":"content 1 "},
			{"title":"title 2","content":"content 2 "},
			{"title":"title 3","content":"content 3 "}
		]
	}`
	var product entities.Product
	json.Unmarshal([]byte(str), &product)
	fmt.Println("Product Info")
	fmt.Println(product.ToString())
	fmt.Println("Category Info")
	fmt.Println("\tCategory Id:", product.Category.Id)
	fmt.Println("\tCategory Name: ", product.Category.Name)
	fmt.Println("Comment List")
	for _, comment := range product.Comments {
		fmt.Println("\t", comment.ToString())
		fmt.Println("\t----------------------------")
	}
}

func main() {
	// Demo1()
	// Demo2()
	// Demo3()
	// Demo4()
	// Demo5()
	Demo6()
}
