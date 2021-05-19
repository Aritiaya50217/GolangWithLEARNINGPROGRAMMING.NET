package product_controller

import (
	"html/template"
	"net/http"
	"strconv"
	"test/entities"
	"test/models"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	// สร้างตัวแปรเก็บข้อมูลลงใน map
	data := map[string]interface{}{
		"products": products,
	}

	tmp, _ := template.ParseFiles("views/product/product.html")
	tmp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("views/product/add.html")
	tmp.Execute(w, nil)
}
func ProcessAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var product entities.Product
	product.Name = r.Form.Get("name")
	// ใช้ ParseFloat เพราะเรากำหนด Price รับข้อมูลประเภท Float
	product.Price, _ = strconv.ParseFloat(r.Form.Get("price"), 64)
	// ใช้ ParseInt เพราะเรากำหนด Quantity รับข้อมูลประเภท Int
	product.Quantity, _ = strconv.ParseInt(r.Form.Get("quantity"), 10, 64)
	product.Description = r.Form.Get("description")

	var productModel models.ProductModel
	// .Create มาจาก func (*ProductModel) Create(product *entities.Product) bool ที่เราสร้างใน product_model.go
	productModel.Create(&product)

	http.Redirect(w, r, "/product", http.StatusSeeOther)
}

// การลบข้อมูลตาม id
func Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var productModel models.ProductModel
	productModel.Delete(id)
	http.Redirect(w, r, "/product", http.StatusSeeOther)

}
// การแก้ไขข้อมูลตาม id
func Edit(w http.ResponseWriter,r*http.Request){
	query := r.URL.Query()
	id,_ := strconv.ParseInt(query.Get("id"),10,64)
	// สร้างตัวแปรที่รับค่ามาจาก ProductModel struct
	var productModel models.ProductModel
	// ค้นหาตาม id
	product,_ := productModel.Find(id)
	data := map[string]interface{}{
		"product" : product,
	}
	tmp, _ := template.ParseFiles("views/product/edit.html")
	tmp.Execute(w, data)
}
