package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var on sync.Once

func main() {
	wg.Add(2)
	go dostuff()
	go dostuff()
	wg.Wait()
	fmt.Println("done")
}

func dostuff() {
	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}

func setup() {
	fmt.Println("setup()")
}
