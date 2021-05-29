package demo1

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello")
}

func hi() {
	fmt.Println("Hi")
}

func Run() {
	go hello()
	go hi()
	// หน่วงเวลา 
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done")
}
