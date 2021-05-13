package demo_controller

import (
	"net/http"
	"testwebsite/entities"
	"text/template"
)

type Data struct {
	Age      int
	Username string
	// นำข้อมูลมาจากโฟล์เดอร์ entities
	// โดยมี type ชื่อ Product
	Product entities.Product
	// สร้างตัวแปร เก็บค่าแบบ map
	Products []entities.Product
}

func Index(w http.ResponseWriter, r *http.Request) {
	// ใส่ที่อยู่ของไฟล์
	tmplt, _ := template.ParseFiles("views/demo_controller/index.html")
	// สร้างตัวแปร เก็บค่าจาก Data
	data := Data{
		Age:      23,
		Username: "abc",
		// เพิ่มค่าได้ทีละชุด
		Product: entities.Product{
			Id:       "p01",
			Name:     "Name01",
			Photo:    "image-16-1024x561.png",
			Price:    1,
			Quantity: 1,
		},
		// สามารถใส่ค่าได้ไม่จำกัดภายใน map[]
		Products: []entities.Product{
			entities.Product{
				Id:       "p01",
				Name:     "Name01",
				Photo:    "image-16-1024x561.png",
				Price:    1,
				Quantity: 1,
			},
			entities.Product{
				Id:       "p02",
				Name:     "Name02",
				Photo:    "image-16-1024x561.png",
				Price:    2,
				Quantity: 2,
			},
			entities.Product{
				Id:       "p03",
				Name:     "Name03",
				Photo:    "image-16-1024x561.png",
				Price:    3,
				Quantity: 3,
			},
		},
	}
	tmplt.Execute(w, data)
}
