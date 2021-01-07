package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	nr              int
	leftCS, rightCS *ChopStick
}

var wg sync.WaitGroup

func main() {
	sticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopStick)
	}

	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{i, sticks[i], sticks[(i+1)%5]}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
	// iets met een Wait
}

// deadlock mogelijkheid
func (p Philosopher) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("philosoher", p.nr, "eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

/* Bovenstaande code kan leiden tot een deadlock. Iedereen heeft alleen z'n linker ChopStick vast.
 * Een mogelijke oplossing kan zijn dat iedereen begint met het oppikken van de CHopstick met het laagste nummer.
 * Voor Philo 4 betekent dat ChopStick 0 (niet 4). Dit kan echter leiden tot starvation. 4 komt nooit meer aan bod.
 * Geen deadlock, wel starvation mogelijk.
 */
