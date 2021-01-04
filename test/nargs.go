package main

import (
	"fmt"
)

func main() {

	d1 := 1

	defer fmt.Println("bye!", d1) // defered values are calculated right away, so 1 is printed
	fmt.Println(getMax(1, 2, 3, 4, 5))

	sli := []int{3, 4, 5, 6, 7}
	fmt.Println(getMax(sli...))
	// so you can either pass 1, 2, 3,... OR a []int slice

	d1++ // di == 2
}

func getMax(values ...int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
