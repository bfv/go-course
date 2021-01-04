package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Hello, playground")
	sli := []int{2, 3, 1}

	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})

	fmt.Println(sli)
}
