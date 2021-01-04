package main

import (
	"fmt"
)

type params struct {
	a, v0, s0 float64
}

func main() {

	p, t := getInput()
	f1 := genDisplaceFn(p.a, p.v0, p.s0)

	fmt.Println("\ndisplacement after", t, "s: ", f1(t))
}

func getInput() (params, float64) {

	var a, v0, s0, t float64

	fmt.Printf("\nenter a, v0 and s0 divided by a blank: ")
	fmt.Scanf("%f %f %f\n", &a, &v0, &s0)

	fmt.Printf("                        enter seconds: ")
	fmt.Scanf("%f\n", &t)

	parms := params{a: a, v0: v0, s0: s0}

	return parms, t
}

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}
