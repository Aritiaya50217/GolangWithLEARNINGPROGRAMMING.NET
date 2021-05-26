package circle

import (
	"math"
	"test/calculate"
)

func Area(r float64) float64 {
	// calculate.Multiply มาจาก โฟลเดอร์ calculate และ Multiply เป็น func 
	return calculate.Multiply(math.Pi,calculate.Multiply(r,r))
}

func Perimeter(r float64) float64 {
	// func Multiply เก็บค่าพารามิเตอร์ 2 ตัว
	return calculate.Multiply(2,calculate.Multiply(math.Pi,r))
}