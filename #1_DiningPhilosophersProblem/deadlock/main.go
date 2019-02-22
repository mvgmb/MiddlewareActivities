package main

import (
	"fmt"
	"time"
)

var forks []byte
var done []bool

func takeFork(fork int) {
	if fork >= len(forks) {
		fork = 0
	}

	for forks[fork] == 1 {
		fmt.Println("Waiting fork", fork)
		time.Sleep(time.Second)
	}

	fmt.Println("Take", fork)
	forks[fork] = 1
}

func eat(philosopher int) {
	fmt.Println("Philosopher", philosopher, "is set")

	time.Sleep(time.Millisecond * 100) // wait's everyone to sit

	takeFork(philosopher + 1) // take right fork
	time.Sleep(time.Second)   // Philosophers are slow
	takeFork(philosopher)     // take left fork

	done[philosopher] = true // Philosopher is done
}

func main() {
	// Prepare table
	noPhilosophers := 5
	forks = make([]byte, noPhilosophers)
	done = make([]bool, noPhilosophers)

	// Philosophers go eat
	for i := 0; i < noPhilosophers; i++ {
		go eat(i)
	}

	// Waits for everyone to finish eating
	var isDone bool
	for !isDone {
		isDone = true
		for i := range done {
			isDone = isDone && done[i]
		}

		time.Sleep(time.Second)
		fmt.Println("Not done")
	}
}
