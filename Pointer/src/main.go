package main

import (
	"fmt"
	"test/entities"
)

func Demo1() {
	var a int = 123
	// ตัวแปร p ชี้ pointer ไปยังที่เก็บของตัวแปร a
	var p *int = &a

	fmt.Println("Address of a:", &a)
	// result 123
	fmt.Println("Value of a:", a)

	fmt.Println("Address of p:", p)
	// result 123 ค่าเท่ากับตัวแปร a
	fmt.Println("Value of p:", *p)
}
func Demo2() {
	a := 456
	// ตัวแปร p ชี้ pointer ไปยังที่เก็บของตัวแปร a
	p := &a
	fmt.Println("Address of a:", &a)
	// result 456
	fmt.Println("Value of a:", a)

	fmt.Println("Address of p:", p)
	// result 456 ค่าเท่ากับตัวแปร a
	fmt.Println("Value of p:", *p)
}

func Change1(a int) {
	a = 2
}

func Change2(p *int) {
	*p = 3
}

func Demo3() {
	a := 1
	Change1(a)
	fmt.Println("a:", a)

	b := 1
	Change2(&b)
	fmt.Println("b:", b)
}

func Swap(a, b int) {
	temp := a
	a = b
	b = temp
}

func Swap2(p, q *int) {
	temp := *p
	*p = *q
	*q = temp

}

func Demo4() {
	a, b := 1, 2
	Swap(a, b)
	fmt.Println("a:", a, "b:", b)

	c, d := 3, 4
	Swap2(&c, &d)
	fmt.Println("c:", c, "d:", d)
}

func Demo5() {
	var product entities.Product
	product.Id = "p1"
	product.Name = "name 1"
	product.Price = 123
	fmt.Println(product)

	var p *entities.Product = &product
	fmt.Println("Product Info")
	fmt.Println("Id:", (*p).Id)
	fmt.Println("Name:", (*p).Name)
	fmt.Println("Price:", (*p).Price)
}
func Demo6() {
	product := entities.Product{
		Id:    "p2",
		Name:  "name 2",
		Price: 222,
	}
	p := &product
	fmt.Println(product)
	fmt.Println("Product Info")
	fmt.Println("Id:", (*p).Id)
	fmt.Println("Name:", (*p).Name)
	fmt.Println("Price:", (*p).Price)
}

// เปลี่ยน  Name
func ChangeProduct1(p *entities.Product) {
	(*p).Name = "def1"
}

func Demo7() {
	product := entities.Product{
		Id:    "p2",
		Name:  "name 2",
		Price: 222,
	}
	ChangeProduct1(&product)
	fmt.Println(product)
}
func main() {
	// Demo1()
	// Demo2()
	// Demo3()
	// Demo4()
	// Demo5()
	// Demo6()
	Demo7()
}
