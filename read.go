package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

// flname (lowercase) -> type is not exported
// Flname (Uppercase) -> type is exported
type flname struct {
	fname string
	lname string
}

const (
	maxStringLength = 20
)

func main() {

	var filename string
	fmt.Println("enter a filename")
	fmt.Scanf("%s", &filename)

	data, _ := ioutil.ReadFile(filename)

	// the slice holding the flname structs
	names := make([]flname, 0)

	s := strings.Split(string(data), "\n")
	for i := range s {

		nameParts := strings.Split(s[i], " ")

		// deal with empty lines
		if len(nameParts) == 1 {
			continue
		}

		name := flname{fname: fixStringLength(nameParts[0]), lname: fixStringLength(nameParts[1])}
		names = append(names, name)
	}

	for i := range names {
		fmt.Println(names[i].fname + " " + names[i].lname)
	}

}

func fixStringLength(s string) string {
	// since len counts bytes we need to check the length via runes
	if utf8.RuneCountInString(s) > maxStringLength {
		runes := []rune(s)
		s = string(runes[:maxStringLength])
	}
	return s
}
