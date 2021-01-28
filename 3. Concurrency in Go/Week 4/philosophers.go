package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Implement the dining Philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each Philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a Philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each Philosopher is numbered, 1 through 5.
When a Philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the Philosopher.
When a Philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the Philosopher.
*/

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	id                            int
	leftChopStick, rightChopStick *ChopStick
}

var eatWGroup sync.WaitGroup

func main() {
	// Assign chop sticks.
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	// Assign philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:             i,
			leftChopStick:  chopSticks[i],
			rightChopStick: chopSticks[(i+1)%5],
		}
		eatWGroup.Add(1)
		go philosophers[i].eat()
	}

	// Wait for all philosophers to eat
	eatWGroup.Wait()
}

func (p Philosopher) eat() {
	// Each philosopher eat 3 times
	for j := 0; j < 3; j++ {
		// Lock the Chop sticks, only 2 can eat at the same time because only have 5 chop sticks and each philosopher need 2.
		// So ⌊5/2⌋ = 2
		p.leftChopStick.Lock()
		p.rightChopStick.Lock()

		fmt.Printf("Philosopher %d is eating\n", p.id+1)
		time.Sleep(time.Second * 3)

		// Unlock
		p.rightChopStick.Unlock()
		p.leftChopStick.Unlock()

		fmt.Printf("Philosopher %d is finished eating\n", p.id+1)
		time.Sleep(time.Second * 3)
	}
	eatWGroup.Done()
}
