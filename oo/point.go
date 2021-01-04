package main

import (
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	x, y float64
}

// p passed by value
func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}

// Scale... p passed as a pointer, dereferencing not necessary (no *p needed)
func (p *Point) Scale(v float64) {
	p.x *= v
	p.y *= v
}

func (p Point) ToString() string {
	return strconv.FormatFloat(p.x, 'f', -1, 64) + ", " + strconv.FormatFloat(p.y, 'f', -1, 64)
}

func main() {
	p1 := Point{x: 4, y: 3}
	p2 := Point{x: 0, y: 0}
	fmt.Println(p1.Distance(p2))

	p1.Scale(2)
	fmt.Println("x, y:", p1.ToString())
}
