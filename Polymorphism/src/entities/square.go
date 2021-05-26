package entities

type Square struct {
	Name string
	A float64
}

func (square Square) Area() float64 {
	return square.A * square.A
}

func (square Square) Perimeter() float64 {
	return square.A*4
}

func (square Square) Type() string {
	return square.Name
}