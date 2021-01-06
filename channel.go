package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const parts = 4

func main() {

	nmbrs, count := getNumbers()
	fmt.Println("unsorted:", nmbrs)

	c := make(chan []int, parts)

	// the 4 (parts) goroutines:
	partSlices := getParts(nmbrs, parts)
	for i := 0; i < parts; i++ {
		go sortArray(partSlices[i], c)
	}

	// now merge it, first copy the first sorted part to the target slice
	sortedAll := make([]int, (count / parts), count)
	copy(sortedAll, <-c)

	for i := 0; i < parts-1; i++ {
		sortedSlice := <-c
		pos := 0
		for j := 0; j < len(sortedSlice); j++ {
			for pos < len(sortedAll) && sortedAll[pos] < sortedSlice[j] {
				pos++
			}
			sortedAll = append(sortedAll[:pos], append([]int{sortedSlice[j]}, sortedAll[pos:]...)...) // insert sortedSlice[j] @ pos
		}
	}

	fmt.Println(sortedAll)
}

func sortArray(arr []int, c chan []int) {
	fmt.Println(arr)
	sort.Ints(arr)
	c <- arr
}

func getNumbers() ([]int, int) {

	var count int

	fmt.Println("To prevent the boring work of entering you can specify the amount of numbers in the initial array.")
	fmt.Printf("Enter the amount number you want sorted: ")
	fmt.Scanf("%d\n", &count)

	// assure random sequence
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	sli := make([]int, count)
	for i := 0; i < count; i++ {
		sli[i] = r1.Intn(count-1) + 1 // + 1 because positive nummbers are required
	}
	return sli, len(sli)
}

// getParts returns partCount slices evenly distributed
func getParts(sli []int, partCount int) [][]int {

	var counter float64
	var start, end int

	result := make([][]int, partCount)
	for i := 0; i < partCount; i++ {
		counter += float64(len(sli)) / float64(partCount)
		end = int(math.Min(math.Ceil(counter), float64(len(sli))))
		result[i] = sli[start:end]
		start = end
	}

	return result
}
