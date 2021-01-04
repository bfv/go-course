package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	foo(s1)
	fmt.Println(s1)
}

func foo(sli []int) {
	sli[0] = sli[0] + 1
}
