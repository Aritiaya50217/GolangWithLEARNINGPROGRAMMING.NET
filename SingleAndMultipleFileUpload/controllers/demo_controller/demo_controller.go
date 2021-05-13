package demo_controller

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmplt, _ := template.ParseFiles("views/demo_controller/index.html")
	tmplt.Execute(w, nil)
}
func SingleUpload(w http.ResponseWriter, r *http.Request) {
	// คำสั่งแสดงไฟล์
	r.ParseMultipartForm(5 * 1024 * 1024)
	// request ข้อมูลมาจาก ฟอร์มที่เราสร้างในไฟล์ .hmtl ตรง name="..."
	file, handler, _ := r.FormFile("file")
	defer file.Close()
	// ชื่อแสดงไฟล์ภาพ และขนาด ใน cmd
	fmt.Println("File Name :", handler.Filename)
	fmt.Println("File Size :", handler.Size)

	// save file ไปยังโฟล์เดอร์ uploads และแสดงชื่อไฟล์
	dst, _ := os.Create("./uploads/" + handler.Filename)
	// หลังจาก save ให้ปิดทันที
	defer dst.Close()
	// copy file
	io.Copy(dst, file)

	tmplt, _ := template.ParseFiles("views/demo_controller/index.html")
	tmplt.Execute(w, nil)
}
func Index2(w http.ResponseWriter, r *http.Request) {
	tmplt, _ := template.ParseFiles("views/demo_controller/index2.html")
	tmplt.Execute(w, nil)
}
func MultipleUpload(w http.ResponseWriter, r *http.Request) {
	// คำสั่งแสดงไฟล์
	r.ParseMultipartForm(5 * 1024 * 1024)
	// request ข้อมูลมาจาก ฟอร์มที่เราสร้างในไฟล์ .hmtl ตรง name="..."
	files := r.MultipartForm.File["files"]
	// len(files) --> การนับจำนวนไฟล์
	fmt.Println("Files : ", len(files))
	// การแสดงไฟล์ตามลำดับ
	for i, fileHander := range files {
		fmt.Println("File : ", i)
		fmt.Println("File Name: ", fileHander.Filename)
		fmt.Println("File Size : ", fileHander.Size)
		// แสดงไฟล์
		file, _ := files[i].Open()
		defer file.Close()
		// Save File
		dst, _ := os.Create("./uploads/" + fileHander.Filename)
		defer dst.Close()

		io.Copy(dst, file)

	}

	tmplt, _ := template.ParseFiles("views/demo_controller/index2.html")
	tmplt.Execute(w, nil)
}
