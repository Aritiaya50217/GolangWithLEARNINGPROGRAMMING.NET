package entities

type Product struct {
	Id,Name string
	Price float64
	Quantity int
	Category Category
}