package main

import (
	"fmt"
	"test/geometry/circle"
	"test/geometry/rectangle"
)

func main() {
	fmt.Println("Circle")
	fmt.Println("Area: ", circle.Area(2))
	fmt.Println("Perimeter: ", circle.Perimeter(2))

	fmt.Println("\nRectangle")
	fmt.Println("Area: ", rectangle.Area(2, 5))
	fmt.Println("Perimeter: ", rectangle.Perimeter(2, 5))
}
