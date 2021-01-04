package main

import (
	"fmt"
	"strings"
)

func main() {

	for {
		var s string
		fmt.Scan(&s)
		s = strings.ToLower(s)
		if strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
