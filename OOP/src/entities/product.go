package entities

import (
	"fmt"
)

type Product struct {
	Id, Name string
	Price    float64
	Quantity int
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id: %s\nName: %s\nPrice: %f\nQuantity: %d",product.Id,product.Name,product.Price,product.Quantity)
}

func (product Product) Total() float64 {
	return product.Price * float64(product.Quantity)
}
