package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}
func (c Cow) Move() {
	fmt.Println("walk")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}
func (b Bird) Move() {
	fmt.Println("fly")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}
func (s Snake) Move() {
	fmt.Println("slither")
}

func main() {
	cowList := make([]Cow, 0, 0)
	birdList := make([]Bird, 0, 0)
	snakeList := make([]Snake, 0, 0)

	for {
		fmt.Printf(">")
		stdScan := bufio.NewScanner(os.Stdin)
		if stdScan.Scan() {
			inputString := strings.Fields(stdScan.Text())
			switch {
			case inputString[0] == "newanimal":
				switch {
				case inputString[2] == "cow":
					var c Cow
					c.name = inputString[1]
					cowList = append(cowList, c)
					fmt.Println("Created it!")
				case inputString[2] == "bird":
					var b Bird
					b.name = inputString[1]
					birdList = append(birdList, b)
					fmt.Println("Created it!")
				case inputString[2] == "snake":
					var s Snake
					s.name = inputString[1]
					snakeList = append(snakeList, s)
					fmt.Println("Created it!")
				}

			case inputString[0] == "query":
				var ani Animal
				if len(cowList) > 0 && inputString[1] != "" {
					for i := 0; i < len(cowList); i++ {
						if cowList[i].name == inputString[1] {
							var c Cow
							c.name = inputString[1]
							ani = c
						}
					}
				}
				if len(birdList) > 0 && inputString[1] != "" {
					for i := 0; i < len(birdList); i++ {
						if birdList[i].name == inputString[1] {
							var b Bird
							b.name = inputString[1]
							ani = b
						}
					}
				}

				if len(snakeList) > 0 && inputString[1] != "" {
					for i := 0; i < len(snakeList); i++ {
						if snakeList[i].name == inputString[1] {
							var s Snake
							s.name = inputString[1]
							ani = s
						}
					}
				}

				if ani != nil {
					switch {
					case inputString[2] == "eat":
						ani.Eat()
					case inputString[2] == "move":
						ani.Move()
					case inputString[2] == "speak":
						ani.Speak()
					default:
						fmt.Println("Input Error!")
					}
				}
			default:
				fmt.Println("Input Error!")
			}
		}
	}

}
