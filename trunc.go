package main

import (
	"fmt"
)

func main() {

	for {
		var f float64
		_, err := fmt.Scan(&f)
		if err != nil {
			break
		}
		fmt.Printf("truncated: %d\n", int(f))
	}
}
