package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct{
	fName string
	lName string
}

func(pct *Name) SetName (fname string, lname string)  {
	pct.fName = fname
	pct.lName = lname
}

func main(){

	var file string
	nameArr := make([]Name,0)
	var name Name

	fmt.Print("Enter the txt file's name (including .txt extension): ")
	fmt.Scan(&file)

	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lineName := strings.Fields(scanner.Text())
		name.SetName(lineName[0], lineName[1])
		nameArr = append(nameArr, name)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	f.Close()

	fmt.Println(nameArr)
}