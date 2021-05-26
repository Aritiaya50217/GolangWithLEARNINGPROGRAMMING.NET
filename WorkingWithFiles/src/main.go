package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"test/entities"
)

func Demo1() {
	// สร้างไฟล์ a.txt
	file, err := os.Create("data/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
		fmt.Println("Done")
	}
	file.Close()
}

func Demo2() {
	file, err := os.Stat("data/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		// แสดงชื่อไฟล์
		fmt.Println("Name:", file.Name())
		// แสดงขนาดไฟล์
		fmt.Println("Size(bytes):", file.Size())
		fmt.Println("Permission:", file.Mode())
		fmt.Printf("Permission: %04o\n", file.Mode().Perm())
		// แสดงเวลา
		fmt.Println("ModTime:", file.ModTime())
	}
}

func Demo3() {
	// ลบไฟล์
	err := os.Remove("data/a.txt")
	if err != nil {
		fmt.Println(err)
	}
}

// การอ่านไฟล์
func Demo4() {
	bytes, err := ioutil.ReadFile("data/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}

func Demo5() {
	file, err := os.Open("data/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var line string = scanner.Text()
			fmt.Println(line)
		}
	}
	file.Close()
}

func Demo6() {
	// สร้างไฟล์
	file, err := os.Create("data/b.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		// เขียนข้อความในไฟล์
		file.WriteString("My File")
	}
	file.Close()
}

func Demo7() {
	// เพิ่มข้อความลงในไฟล์
	file, err := os.OpenFile("data/b.txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(file, "Line 3")
	}
	file.Close()
}
func Demo8() {
	// open file
	file, err := os.Open("data/products.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		var products []entities.Product
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var line string = scanner.Text()
			result := strings.Split(line, ",")
			price, _ := strconv.ParseFloat(result[2], 64)
			quantity, _ := strconv.ParseInt(result[3], 10, 64)
			products = append(products, entities.Product{
				// [0] ตำแหน่งแรก
				Id:   result[0],
				Name: result[1],
				// ไม่ต้องใส่ [] เพราะสร้างตัวแปรไว้ด้านบนแล้ว
				Price:    price,
				Quantity: quantity,
			})
		}
		fmt.Println("product List form CSV File")
		for _, product := range products {
			fmt.Println(product.ToString())
			fmt.Println("----------------------------")
		}
	}
	file.Close()
}

func main() {
	// Demo1()
	// Demo2()
	// Demo3()
	// Demo4()
	// Demo5()
	// Demo6()
	// Demo7()
	Demo8()
}
