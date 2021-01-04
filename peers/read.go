package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxLength = 20
)

type Name struct {
	fname string
	lname string
}

var fileName string

func main() {

	namesList := make([]Name, 0, 1)

	fmt.Print("enter the filename : ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fullName := strings.Split(line, " ")
		var p Name

		//Duplicated code : After learning functions with i'll make this logic reusable
		fmt.Println("taille : ", len(fullName[0]))
		if len(fullName[0]) > maxLength {
			p.fname = fullName[0][:maxLength]
		} else {
			p.fname = fullName[0]
		}

		if len(fullName[1]) > maxLength {
			p.lname = fullName[1][:maxLength]
		} else {
			p.lname = fullName[1]
		}

		namesList = append(namesList, p)
	}

	fmt.Println("\nList of names\n===============")
	for _, v := range namesList {
		fmt.Printf("%s %s\n", v.fname, v.lname)
	}

	file.Close()
}
