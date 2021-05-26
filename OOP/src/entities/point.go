package entities

import (
	"math"
)
type Point struct {
	X,Y int 
}

func (point Point) Lenght(b Point)float64 {
	// แปลง type เพราะ ตัวแปร b เป็น float64
	return math.Sqrt(float64((point.X-b.X) *(point.X - b.X) + (point.Y-b.Y) * (point.Y-b.Y)))
}