package main

import (
	"fmt"
	"test/entities"
)

func Demo1() {

	circle1 := entities.Circle{"Circle 1", 5}
	circle2 := entities.Circle{"Circle 2", 2}

	rectangle1 := entities.Rectangle{"Retangle 1", 2, 4}
	rectangle2 := entities.Rectangle{"Retangle 1", 7, 3}

	square1 := entities.Square{"Square 1", 5}
	square2 := entities.Square{"Square 2", 9}

	geometries := []entities.Geometry{circle1, square1, rectangle1, circle2, square2, rectangle2}
	for _, geometry := range geometries {
		fmt.Println(geometry.Type())
		fmt.Println("Area:", geometry.Area())
		fmt.Println("Perimeter:", geometry.Perimeter())
		fmt.Println("--------------------")
	}
}
func Demo2() {
	dog1 := entities.Dog{"Dog 1"}
	dog2 := entities.Dog{"Dog 2"}

	cat1 := entities.Cat{"Cat 1"}
	cat2 := entities.Cat{"Cat 2"}

	duck1 := entities.Duck{"Duck 1"}
	duck2 := entities.Duck{"Duck 2"}

	animals := []entities.Animal{dog1, duck1, cat1, dog2, duck2, cat2}
	for _, animal := range animals {
		fmt.Println(animal.Type())
		fmt.Println(animal.Speak())
		fmt.Println("--------------------")
	}
}

func main() {
	// Demo1()
	Demo2()
}
