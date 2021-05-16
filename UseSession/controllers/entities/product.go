package entities

type Product struct {
	// แสดงข้อมูลในรูปแบบของ json
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Status   bool    `json:"status"`
}
