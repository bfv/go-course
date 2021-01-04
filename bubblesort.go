package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := getInput()
	bubbleSort(numbers)

	fmt.Println("ordered:", numbers)

}

func bubbleSort(numbers []int) {

	ordered := false
	for !ordered {
		ordered = true
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				swap(numbers, i)
				ordered = false
			}
		}
	}
}

func swap(numbers []int, i int) {
	t := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = t
}

func getInput() []int {

	numbers := make([]int, 0, 10)

	fmt.Printf("enter a numbers (non number to stop):\n")

	i := 0
	for i < 10 {
		var s string
		fmt.Scan(&s)

		v, err := strconv.Atoi(s)
		if err != nil {
			break
		}

		numbers = append(numbers, v)
		i++
	}

	fmt.Println("input:  ", numbers)

	return numbers
}
