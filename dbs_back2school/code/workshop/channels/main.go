package main

import (
	"fmt"
	"time"
)

func SendReceiveExample() {
	// EXAMPLE START OMIT

	greenPipe := make(chan string)
	greenPipe <- "Mario" // Send a value to a channel
	mario := <-greenPipe // Receive a value from a channel

	// EXAMPLE END OMIT
	fmt.Println(mario)
}

// DELAY START OMIT

func deplayedPrint(i int, valChan chan<- int) {
	time.Sleep(time.Millisecond * 500)
	valChan <- i
}

// DELAY END OMIT

func main() {
	start := time.Now()
	// MAIN START OMIT
	ch := make(chan int)
	for i := 1; i < 6; i++ {
		go deplayedPrint(i, ch)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	// MAIN END OMIT
	fmt.Printf("Program took %s\n", time.Since(start))
}
