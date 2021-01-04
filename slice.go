package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var s string
	// The previous editor noticed that the output has zero's which he didn't enter
	// and "blamed" the way I created my slice. However, the instructions are quite clear
	// that the length (not the capacity) should be 3. So I think my solution is OK (in that respect)
	input := make([]int, 3)
	var count int

	for {
		fmt.Print("Please enter a number or X to finish : ")
		fmt.Scan(&s)
		if strings.ToUpper(s) == "X" {
			break
		}

		v, err := strconv.Atoi(s)
		if err == nil {
			input = append(input, v)
			sort.Slice(input, func(i, j int) bool {
				return input[i] < input[j]
			})
			fmt.Println(input)
			count++
		}
	}
}
