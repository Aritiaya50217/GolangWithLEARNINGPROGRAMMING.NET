package entities

import (
	"fmt"
)
type Product struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Status   bool    `json:"status"`
	// ดึงค่ามาจาก Category struct
	Category Category `json:"category"`
	// ดึงค่ามาจาก Comment struct
	Comments []Comment `json:"comments"`
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id:%s\nName:%s\nPrice:%f\nQuantity:%d\nStatus:%t",product.Id,product.Name,product.Price,product.Quantity,product.Status)
}
