package main

import (
	"fmt"
	"sort"
	"strings"
	"test/entities"
)

func Demo1() {
	// เก็บข้อมูล 3 ชุด
	var products [3]entities.Product
	products[0] = entities.Product{
		Id:       "p1",
		Name:     "name 1",
		Price:    3.9,
		Quantity: 6,
		Category: entities.Category{
			Id:   "c1",
			Name: "Category 1",
		},
	}
	products[1] = entities.Product{
		Id:       "p2",
		Name:     "name 2",
		Price:    2.9,
		Quantity: 4,
		Category: entities.Category{
			Id:   "c2",
			Name: "Category 2",
		},
	}
	products[2] = entities.Product{
		Id:       "p3",
		Name:     "name 3",
		Price:    2.9,
		Quantity: 4,
		Category: entities.Category{
			Id:   "c3",
			Name: "Category 3",
		},
	}
	fmt.Println(products)
	fmt.Println("Size:", len(products))
	fmt.Println("Product List 1")
	for i := 0; i < len(products); i++ {
		fmt.Println("Product ", i)
		Display(products[i])
		fmt.Println("**************************")
	}
	fmt.Println("Product List 2")
	for i, product := range products {
		fmt.Println("Product ", i)
		Display(product)
		fmt.Println("**************************")
	}
}

func Display(product entities.Product) {

	fmt.Println("Id", product.Id)
	fmt.Println("Name", product.Name)
	fmt.Println("Price", product.Price)
	fmt.Println("Quantity", product.Quantity)

	fmt.Println("Category Id", product.Category.Id)
	fmt.Println("Category Name", product.Category.Name)

}
func Demo2() {
	// เอาข้อมูลมาจาก func FindAll()
	products := FindAll()
	fmt.Print("Price ASC")
	SortPriceASC(products)
	fmt.Println(products)

	fmt.Println("Price DESC")
	SortPriceDESC(products)
	fmt.Println(products)
}

func Demo3() {
	products := FindAll()
	max, min := MaxAndMin(products)
	fmt.Println("Product Max Info")
	Display(max)

	fmt.Println("Product Min Info")
	Display(min)
}
func Demo4() {
	keyword := "note"
	products := Search(FindAll(), keyword)
	fmt.Println(products)

}

func Demo5() {
	s := Total(FindAll())
	fmt.Println("Total:", s)
}

func Demo6() {

	result1 := Exists("p1", FindAll())
	fmt.Println("result 1:", result1)
	result2 := Exists("p2", FindAll())
	fmt.Println("result 2:", result2)
	result3 := Exists("p3", FindAll())
	fmt.Println("result 3:", result3)

	// ลองใส่ id ที่ไม่มีอยู่
	result4 := Exists("p4", FindAll())
	fmt.Println("result 4:", result4)
}

func Exists(id string, products []entities.Product) bool {
	for _, product := range products {
		// ถ้า product.Id ที่เราสร้าง เท่ากับ ค่า id ที่ใส่เข้าไป
		if product.Id == id {
			return true
		}
	}
	return false
}

// คำนวณราคา
func Total(products []entities.Product) (s float64) {
	s = 0
	for _, product := range products {
		// แปลง type
		s += product.Price * float64(product.Quantity)
	}
	return
}

// ค้นหา product.Name
func Search(products []entities.Product, keyword string) (result []entities.Product) {
	for _, product := range products {
		// แปลงให้เป็นพิมพ์เล็กทั้ง product.Name และ keyword
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(keyword)) {
			result = append(result, product)
		}
	}
	return
}

// เรียงลำดับตามราคา (price) จากน้อยไปมาก
func SortPriceASC(products []entities.Product) {
	sort.Slice(products, func(i, j int) bool {
		// แสดงค่า Price ที่น้อยสุดก่อน
		return products[i].Price <= products[j].Price
	})
}

// เรียงลำดับตามราคา (price) จากมากไปน้อย
func SortPriceDESC(products []entities.Product) {
	sort.Slice(products, func(i, j int) bool {
		// แสดงค่า Price ที่น้อยสุดก่อน
		return products[i].Price >= products[j].Price
	})
}

func MaxAndMin(products []entities.Product) (max entities.Product, min entities.Product) {
	max = products[0]
	min = products[0]
	for _, product := range products {
		// Find Max
		if max.Price <= product.Price {
			max = product
		}
		// Find Min
		if min.Price >= product.Price {
			min = product
		}
	}
	return
}

func FindAll() (products []entities.Product) {
	// ไม่ระบุว่าเก็บข้อมูลกี่ชุด (slice)
	products = []entities.Product{
		{
			Id:       "p1",
			Name:     "Laptop 1",
			Price:    1,
			Quantity: 6,
			Category: entities.Category{
				Id:   "c1",
				Name: "Category 1",
			},
		},
		{
			Id:       "p2",
			Name:     "NoteBook 2",
			Price:    2,
			Quantity: 4,
			Category: entities.Category{
				Id:   "c2",
				Name: "Category 2",
			},
		},
		{
			Id:       "p3",
			Name:     "Computer 3",
			Price:    3,
			Quantity: 4,
			Category: entities.Category{
				Id:   "c3",
				Name: "Category 3",
			},
		},
	}
	return
}

func main() {
	// Demo2()
	// Demo3()
	// Demo4()
	// Demo5()
	Demo6()
}
