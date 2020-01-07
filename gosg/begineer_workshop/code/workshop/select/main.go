package main

import (
	"fmt"
	"math/rand"
)

// FLIP START OMIT
func flipACoin(heads chan<- int, tails chan<- int) {
	if rand.Float32() < 0.5 {
		heads <- 1
	} else {
		tails <- 1
	}
}

// FLIP END OMIT

func main() {
	heads := make(chan int)
	tails := make(chan int)
	// MAIN START OMIT
	for i := 0; i < 10; i++ {
		go flipACoin(heads, tails)
	}
	for i := 0; i < 10; i++ {
		select {
		case <-heads:
			fmt.Println("Heads!")
		case <-tails:
			fmt.Println("Tails!")
		}
	}
	// MAIN END OMIT
}
