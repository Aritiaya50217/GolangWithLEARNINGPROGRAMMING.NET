package main

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello")
}

func Hi(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

func Add(a, b int) int {
	return a + b
}
func Calculate(a, b int) (int, int, int, float64) {
	result1 := a + b
	result2 := a - b
	result3 := a * b
	// แปลง type ให้เป็น float64 จึงจะหารได้
	result4 := float64(a) / float64(b)

	return result1, result2, result3, result4
}

func Calculate2(a, b int) (result1 int, result2 int, result3 int, result4 float64) {
	result1 = a + b
	result2 = a - b
	result3 = a * b
	// แปลง type ให้เป็น float64 จึงจะหารได้
	result4 = float64(a) / float64(b)
	return
}

// ไม่จำกัด parameters เก็บค่าเป็น int
func Sum(parameters ...int) int {
	// len การนับ
	fmt.Println("Size:", len(parameters))
	s := 0
	// v นับตำแหน่งหรือลำดับใน ภายใน parameters
	for _, v := range parameters {
		// เอาแต่ละตำแหน่งมา + กัน
		s += v
	}
	return s
}

func Find(x int, numbers ...int) bool {
	// v นับตำแหน่งหรือลำดับใน ภายใน numbers
	for _, v := range numbers {
		// ถ้า x เป็นเลขเดียวกันกับ v
		if x == v {
			return true
		}
	}
	return false

}

func main() {
	Hello()
	// "ABC" คือ name
	fmt.Println(Hi("ABC"))
	fmt.Println("Result:", Add(2, 3))

	r1, r2, r3, r4 := Calculate(10, 4)
	fmt.Println("10 + 4 = ", r1)
	fmt.Println("10 - 4 = ", r2)
	fmt.Println("10 * 4 = ", r3)
	fmt.Println("10 / 4 = ", r4)

	re1, re2, re3, re4 := Calculate2(10, 4)
	fmt.Println("10 + 4 = ", re1)
	fmt.Println("10 - 4 = ", re2)
	fmt.Println("10 * 4 = ", re3)
	fmt.Println("10 / 4 = ", re4)

	fmt.Println("Sum 1:", Sum(1, 2))
	fmt.Println("Sum 2:", Sum(10, 2, -6, 3))
	fmt.Println("Sum 3:", Sum(-7, 3, 8))

	fmt.Println("Find1:", Find(4, 5, 2, 1, -3))
	fmt.Println("Find2:", Find(4, 5, 2, 1, -30, 4))

}
