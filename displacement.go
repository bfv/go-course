package main

import (
	"fmt"
	"strconv"
)

type params struct {
	a, v0, s0 float64
}

func main() {

	p1, t := getInput()
	f1 := genDisplaceFn(p1.a, p1.v0, p1.s0)

	fmt.Println("\ndisplacement after", t, "s: ", f1(t))
}

func getInput() (params, float64) {

	var sa, sv0, ss0, st string

	fmt.Printf("\nenter a, v0 and s0 divided by a blank: ")
	fmt.Scanf("%s %s %s\n", &sa, &sv0, &ss0)

	fmt.Printf("                        enter seconds: ")
	fmt.Scanf("%s\n", &st)

	a, _ := strconv.ParseFloat(sa, 64)
	v0, _ := strconv.ParseFloat(sv0, 64)
	s0, _ := strconv.ParseFloat(ss0, 64)
	parms := params{a: a, v0: v0, s0: s0}

	t, _ := strconv.ParseFloat(st, 64)

	return parms, t
}

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}
