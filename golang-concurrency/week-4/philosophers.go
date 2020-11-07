package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id, count int
	leftChop, rightChop *Chopstick
}

func (p Philosopher) Eat(wg *sync.WaitGroup, c chan *Philosopher) {
	for i := 0; i < 3; i++ {
		c <- &p

		if p.count < 3 {
			p.leftChop.Lock()
			p.rightChop.Lock()
			fmt.Println("starting to eat", p.id)

			p.count++

			fmt.Println("finishing eating", p.id)
			p.rightChop.Unlock()
			p.leftChop.Unlock()
			wg.Done()
		}
	}
}

func host(wg *sync.WaitGroup, c chan *Philosopher) {
	for {
		if len(c) == 2 {
			<- c
			<- c

			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	chopsticks, philosophers, c := make([]*Chopstick, 5), make([]*Philosopher, 5), make(chan *Philosopher, 2)

	wg.Add(15)

	for i := 0; i < 5; i++ {
		chopsticks[i] = &Chopstick{}
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, 0, chopsticks[i], chopsticks[(i + 1) % 5]}
	}

	go host(&wg, c)

	for _, phi := range philosophers {
		go phi.Eat(&wg, c)
	}

	wg.Wait()
}
