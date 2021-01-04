package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64

	// Prompt Values
	fmt.Printf("Please input acceleration: ")
	fmt.Scan(&a)
	fmt.Printf("Please input initial velocity: ")
	fmt.Scan(&v0)
	fmt.Printf("Please input initial displacement: ")
	fmt.Scan(&s0)
	fmt.Printf("Please input time: ")
	fmt.Scan(&t)

	// Generate Displacement Function
	myGenFn := GenDisplaceFn(a, v0, s0)

	// Compute Displacement
	fmt.Printf("%f", myGenFn(t))
}
