package entities

import (
	"math"
)

type Circle struct {
	Name string
	R    float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.R * circle.R
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.R
}

func (circle Circle) Type() string {
	return circle.Name
}
