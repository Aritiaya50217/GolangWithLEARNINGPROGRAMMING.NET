package demo3

import (
	"fmt"
	"time"
)

func display_numbers() {
	// วนลูปแสดงเลขตั้งแต่ 1-10
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func display_characters() {
	// วนลูปแสดงตัวอักษร a - e
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Printf(" %c ", ch)
		time.Sleep(200 * time.Millisecond)
	}
}

func Run() {
	// ผลลัพธ์ที่ได้จะแสดงตัวเลข 1 -10 กับอักษร a - e ปนกัน เพราะ time.Sleep ที่เราตั้งไว้มีเวลาห่างกันเล็กน้อย
	go display_numbers()
	go display_characters()
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Done")

}
