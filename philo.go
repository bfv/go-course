package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	nr              int
	leftCS, rightCS *ChopStick
}

const philoCount = 5
const mealCount = 3

var wg sync.WaitGroup
var onceDo sync.Once

func main() {

	ca := make(chan bool, 2) // available channel
	var cr chan bool

	sticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopStick)
	}

	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{i + 1, sticks[i], sticks[(i+1)%5]}
	}

	for i := 0; i < philoCount; i++ {
		wg.Add(1)
		go philos[i].eat(ca, cr)
	}

	go host(ca, cr)

	wg.Wait()
	// iets met een Wait
}

// deadlock mogelijkheid
func (p Philosopher) eat(ca chan bool, cr chan bool) {

	for i := 0; i < 3; i++ {

		fmt.Println("philosopher", p.nr, " waiting for meal", i)
		<-ca // wait for availability
		fmt.Printf("p=%d,eat=%d locking\n", p.nr, i)
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("philosopher", p.nr, " start eating", i)
		time.Sleep(time.Second)
		fmt.Println("philosopher", p.nr, " end eating", i)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		fmt.Printf("p=%d,eat=%d unlocked again\n", p.nr, i)
		cr <- true // signal readiness
	}
	wg.Done()
}

func host(ca chan bool, cr chan bool) {

	totalMeals := philoCount * mealCount

	fmt.Println("start host")
	ca <- true
	fmt.Println("available 0")
	ca <- true
	fmt.Println("available 1")

	for i := 0; i < totalMeals-2; i++ {
		fmt.Println("host waits for ready of ", i)
		<-cr // wait for ready signal
		fmt.Println("ready received", i)
		ca <- true
		fmt.Println("available", i+2)
	}
}
