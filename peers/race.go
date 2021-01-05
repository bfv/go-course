package main

import (
	"fmt"
	"time"
)

//adding of iteration no signifies that the increment operation of x happens sometimes multiple times before the statement in the same function is printed

// race condition is when the same data is accessed by two different operations which are running concurrently hence providing a non deterministic end result

func PrintX(x *int, i int) {
	*x += 2
	fmt.Printf("conc : %d of %d iteration \n", *x, i)
}
func main() {

	fmt.Printf("every time you execute the compiled code there will be a different output of x hence providing with instance of race condition\n")

	x := 1
	x = x + 1
	for i := 0; i < 5; i++ {
		go PrintX(&x, i)
		fmt.Printf("main : %d  \n", x)
	}

	time.Sleep(time.Second) //added sleep to let all the function finish execution and not only the main function.

}
