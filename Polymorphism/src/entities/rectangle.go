package entities


type Rectangle struct {
	Name string
	A float64
	B float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.A * rectangle.B
}

func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.A + rectangle.B) *2
}

func (rectangle Rectangle) Type() string {
	return rectangle.Name
}