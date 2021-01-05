package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// a util function for displaying prompt and scan for user input
func promptAndScan(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	name := promptAndScan("enter your name: ")
	address := promptAndScan("enter your address: ")

	// build a map
	addrByName := map[string]string{
		"name":    name,
		"address": address}

	fmt.Printf("\n ---- output ---\n")

	// convert to json and print
	barr, _ := json.Marshal(addrByName)
	fmt.Println("json bytes: ")
	fmt.Println(barr)

	fmt.Println("json string: ")
	fmt.Println(string(barr))
}
