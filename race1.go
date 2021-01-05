package main

import (
	"fmt"
	"sync"
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

	var wg sync.WaitGroup

	wg.Add(2)
	go f("t1", &wg)
	time.Sleep(time.Millisecond)
	go f("t2", &wg)

	wg.Wait()
	fmt.Println("end")
}

func f(name string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		x++
		fmt.Printf("%s %d\n", name, x)
	}
	wg.Done()
}
