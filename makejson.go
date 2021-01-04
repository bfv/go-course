package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string

	na := make(map[string]string)

	fmt.Printf("enter name: ")
	fmt.Scanf("%s\n", &name)
	na["name"] = name

	fmt.Printf("enter address: ")
	fmt.Scanf("%s\n", &addr)
	na["address"] = addr

	byteArray, _ := json.Marshal(na)

	fmt.Println(string(byteArray))
}
