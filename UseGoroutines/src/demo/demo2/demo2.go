package demo2

import (
	"fmt"
	"time"
)

func Process1(){
	fmt.Println("Start Process 1")
	// หน่วงเวลา 
	time.Sleep(100 * time.Millisecond)
	fmt.Println("End Process 1")
}

func Process2(){
	fmt.Println("Start Process 2")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("End Process 2")
}

func Process3(){
	fmt.Println("Start Process 3")
	time.Sleep(150 * time.Millisecond)
	fmt.Println("End Process 3")
}

func Run() {
	go Process1()
	go Process2()
	go Process3()
	// หน่วงเวลา
	time.Sleep( 500 * time.Millisecond)
	fmt.Println("Done")
}