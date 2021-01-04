package main

import (
	"fmt"
	"math"
)

func main() {

	f := incFn

	var f1 func(int) int

	f1 = incFn

	fmt.Println(f(2))
	fmt.Println(f1(5))

	fmt.Println(applyIt(f1, 8)) // pass f1 to applyIt function

	distFromZero := makeDistOrigin(0, 0)
	fmt.Println("length: ", distFromZero(3, 4))

	dist2 := makeDistOrigin(3, 3)
	fmt.Println("length2: ", dist2(6, 7))
}

func incFn(x int) int { return x + 1 }

// applyIt takes a function (f) as a parameter with an integer params and int return value
func applyIt(f func(int) int, x int) int {
	return f(x)
}

func makeDistOrigin(originX, originY float64) func(float64, float64) float64 {
	f := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-originX, 2) + math.Pow(y-originY, 2))
	}
	return f
}
