package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const count = 5
const maxEatCount = 3
const maxConcurrent = 2

var concurrentGroup *sync.WaitGroup
var wg sync.WaitGroup

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id              int
	eatCount        int
	leftCS, rightCS *Chopstick
}

func (p *Philosopher) Eat() {
	randomPhilo := rand.Float64()
	if randomPhilo < 0.5 {
		p.leftCS.Lock()
		p.rightCS.Lock()
	} else {
		p.rightCS.Lock()
		p.leftCS.Lock()
	}

	fmt.Printf("starting to eat %v \n", p.id)
	p.eatCount++
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("finishing eating %v \n", p.id)
	p.leftCS.Unlock()
	p.rightCS.Unlock()
	concurrentGroup.Done()
	wg.Done()
}

func (p *Philosopher) EatControlPerPhilo(philoChan chan *Philosopher) {
	for eat := 0; eat < maxEatCount; eat++ {
		philoChan <- p
	}
}

type Host struct {
	concurrentExecutions int
}

func (h Host) eatControl(philoChan chan *Philosopher) {
	for {
		philo := <-philoChan
		if h.concurrentExecutions == maxConcurrent {
			concurrentGroup.Wait()
			h.concurrentExecutions = 0
		}
		concurrentGroup.Add(1)
		h.concurrentExecutions++
		go philo.Eat()
	}
}

func EveryoneEat(philosophers []Philosopher, philoChan chan *Philosopher) {
	for i := 0; i < count; i++ {
		go philosophers[i].EatControlPerPhilo(philoChan)
	}
}

func main() {
	concurrentGroup = &sync.WaitGroup{}
	wg = sync.WaitGroup{}

	host := Host{}
	host.concurrentExecutions = 0

	ChopSticks := make([]Chopstick, count)
	for i := 0; i < count; i++ {
		ChopSticks[i] = Chopstick{}
	}
	Philosophers := make([]Philosopher, count)
	for i := 0; i < count; i++ {
		Philosophers[i] = Philosopher{leftCS: &ChopSticks[i], rightCS: &ChopSticks[(i+1)%count], id: i + 1}
	}

	philoChan := make(chan *Philosopher)

	go host.eatControl(philoChan)
	wg.Add(count * maxEatCount)
	EveryoneEat(Philosophers, philoChan)
	wg.Wait()
}
