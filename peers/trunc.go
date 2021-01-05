package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	fmt.Println("Please enter a floating point number and press enter")
	reader := bufio.NewReader(os.Stdin)
	floatingNum, err := reader.ReadString('\n')
	truncatedNum := strings.Split(floatingNum, ".")[0]
	if err != nil {
		fmt.Println("Error while reading string", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Not a valid floating number", err)
		os.Exit(1)
	}
	fmt.Println("Truncated number - ", truncatedNum)
}
