package main
 import "fmt"

 type Animal struct {
 	food string
 	locomotion string
 	noise string
 }

 func (an *Animal) Eat() {
 	fmt.Println(an.food)
 }
  func (an *Animal) Move() {
 	fmt.Println(an.locomotion)
 }
  func (an *Animal) Speak() {
 	fmt.Println(an.noise)
 }

func main(){
	var mapanimal map[string]Animal
	mapanimal = make(map[string]Animal)
mapanimal["cow"] = Animal{food : "grass",locomotion : "walk",noise : "moo"}
mapanimal["bird"] = Animal{food : "worms",locomotion : "fly",noise : "peep"}
mapanimal["snake"] =  Animal{food : "mice",locomotion : "slither",noise : "hss"}

 for {
 	var a, info string
 	fmt.Printf(">")
	fmt.Scan(&a,&info)

	anim := Animal(mapanimal[a])
	switch {
		case info == "eat" :
		anim.Eat()
		case info == "move" :
		anim.Move()
		case info == "speak" :
		anim.Speak()
		default:
	}

 }
}