package main

import (
	"fmt"
	"test/entities"
)

func Demo1() {
	var product1 entities.Product
	product1.Id = "p1"
	product1.Name = "Name 1"
	product1.Price = 4.5
	product1.Quantity = 2
	fmt.Println(product1.ToString())
	fmt.Println("Total:", product1.Total())
}

func Demo2() {
	product2 := entities.Product{
		Id:       "p2",
		Name:     "name 2",
		Price:    4,
		Quantity: 5,
	}
	fmt.Println(product2.ToString())
	fmt.Println("Total:", product2.Total())

}

func Demo3() {

	a := entities.Point{X: 4, Y: 7}
	b := entities.Point{X: 1, Y: 3}
	c := entities.Point{X: 2, Y: 4}
	perimeter := a.Lenght(b) + a.Lenght(c) + b.Lenght(c)
	fmt.Println("Perimeter:", perimeter)
}
func Demo4() {
	products := []entities.Product{
		entities.Product{
			Id:       "p1",
			Name:     "name 1",
			Price:    2.6,
			Quantity: 3,
		},
		entities.Product{
			Id:       "p2",
			Name:     "name 2",
			Price:    4.6,
			Quantity: 4,
		},
		entities.Product{
			Id:       "p3",
			Name:     "name 3",
			Price:    1.6,
			Quantity: 8,
		},
	}
	var total float64 = 0
	for _, product := range products {
		// func Total คือ product.Price * float64(product.Quantity)
		total += product.Total()
	}
	fmt.Println("Total:", total)

}

func main() {
	// Demo1()
	// Demo2()
	// Demo3()
	Demo4()
}
