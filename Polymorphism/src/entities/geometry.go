package entities

type Geometry interface {
	Area() float64
	Perimeter() float64
	Type() string
}