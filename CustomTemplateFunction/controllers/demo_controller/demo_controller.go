package demo_controller

import (
	"html/template"
	"net/http"
	"strings"
	"test/entities"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"fullName": "Oil",
		// นำข้อมูลมาจาก Product struct
		"product": entities.Product{
			Id:       "P01",
			Name:     "name 1 ",
			Price:    3.3,
			Quantity: 5,
			Status:   true,
		},
	}
	tmp, _ := template.New("index.html").Funcs(template.FuncMap{
		// แปลงข้อความให้เป็นพิมพ์ใหญ่
		"upper": func(name string) string {
			// return ออกมาเป็นตัวพิมพ์ใหญ่ ผ่านตัวแปร name
			return strings.ToUpper(name)
		},
		// คำสั่งการแสดง status
		"displayStatus": func(status bool) string {
			if status {
				return "Show"
			}
			return "Hide"
		},
		// คำสั่งการแสดงราคารวมทั้งหมด
		"total": func(product entities.Product) float64 {
			// return ผลลัพธ์ที่ได้จากการเอาราคา * จำนวนสินค้า
			// ต้องแปลงให้ type เป็นชนิดเดียวกัน ในที่นี้แปลง Quantity ให้เป็น float64
			return product.Price * float64(product.Quantity)

		},
		// ตรวจสอบจำนวน อ้างอิงข้อมูลจาก Product struct ที่นำมา map[] ด้านบน
		"checkStock": func(product entities.Product) string {
			if product.Quantity == 0 {
				return "Out of Stock"
			}
			return "In Stock"
		},
	}).ParseFiles("views/demo_controller/index.html")
	tmp.Execute(w, data)
}
