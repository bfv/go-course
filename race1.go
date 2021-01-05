package main

import (
	"fmt"
	"time"
)

/* function f iterates 5 times, incrementing x which is shared
 * between the two go routines, f("t1") & f("t2")
 * The race condition is in that it becomes unpredictable which
 * go routine prints what, since they both alter the context (var x)
 */

var x int

func main() {

	fmt.Println("start")
	go f("t1")
	// fmt.Println("hi!")
	go f("t2")

	time.Sleep(time.Second) // making sure everything can finish
	fmt.Println("end")
}

func f(name string) {
	for i := 0; i < 5; i++ {
		x++
		fmt.Printf("%s %d\n", name, x)
	}
}
