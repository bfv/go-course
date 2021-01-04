package main

import "fmt"

// Animal struct holds the values for eating, movement and "speaking
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints food
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

// Move prints locomotion
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

// Speak prints the noise
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

// Animals is the Map which the animal name to the appropriate struct
type Animals map[string]Animal

func main() {

	var animal, action string

	animals := getAnimals()

	fmt.Println("Enter an animal (cow, bird or snake) and an action (eat, move or speak), type exit to stop")

	proceed := true
	for proceed {

		fmt.Printf("> ")
		fmt.Scanf("%s %s\n", &animal, &action)

		if animal == "exit" {
			return
		}

		chosenAnimal := animals[animal]
		switch action {
		case "eat":
			chosenAnimal.Eat()
		case "move":
			chosenAnimal.Move()
		case "speak":
			chosenAnimal.Speak()
		}
	}
}

func getAnimals() Animals {

	animals := make(map[string]Animal)

	animals["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	return animals
}
