package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2))
}

func main() {
	p1 := NewPoint(2, 5)
	p2 := NewPoint(8, 10)
	fmt.Printf("Расториние между точками с координатами p1:x%dy%d и p2:x%dy%d, - %f", p1.x, p1.y, p2.x, p2.y, Distance(*p1, *p2))
}
