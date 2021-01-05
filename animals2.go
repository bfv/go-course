package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}
func (cow Cow) Move() {
	fmt.Println("walk")
}
func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}
func (bird Bird) Move() {
	fmt.Println("fly")
}
func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}
func (snake Snake) Move() {
	fmt.Println("slither")
}
func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	var command, name, sub string

	animals := make(map[string]Animal)

	fmt.Println("type newanimal <name> cow|bird|snake OR query <name> eat|speak|move, exit to exit")

	for {

		fmt.Printf("> ")
		fmt.Scanf("%s %s %s\n", &command, &name, &sub)

		if command == "exit" {
			return
		}

		var animal Animal
		if command == "newanimal" {
			switch sub {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			}
			animals[name] = animal
			fmt.Println("Created it!")
		} else if command == "query" {
			animal = animals[name]
			switch sub {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		} else {
			fmt.Println("<syntax-error>")
		}
	}
}
