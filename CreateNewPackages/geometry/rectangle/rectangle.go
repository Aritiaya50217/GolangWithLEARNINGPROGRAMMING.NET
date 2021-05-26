package rectangle

import (
	"test/calculate"
)

func Area(a,b float64 ) float64 {
	return calculate.Multiply(a,b)
}

func Perimeter(a,b float64) float64 {
	return calculate.Multiply(calculate.Add(a,b),2)
}