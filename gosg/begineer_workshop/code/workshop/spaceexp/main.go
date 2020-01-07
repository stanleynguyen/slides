package main

import (
	"fmt"
	"math/rand"
	"time"
)

// EXERCISE START OMIT

type Astronaut struct {
	Name string
	Ship chan string
}

func main() {
	theHermes := make(chan string)
	neiArmstrong := Astronaut{"NEIL ARMSTRONG", theHermes}
	sallyRide := Astronaut{"SALLY RIDE", theHermes}
	done := make(chan bool)
	go neiArmstrong.Explore(done)
	go sallyRide.Explore(done)
	theHermes <- "start"
	<-done
}

// SOLUTION START OMIT
func (a *Astronaut) Explore(done chan bool) {
	// EXERCISE END OMIT
	for {
		msg := <-a.Ship
		if msg == "GEM" {
			fmt.Println("Mission success!")
			done <- true
			break
		} else {
			time.Sleep(1000 * time.Millisecond) // delay for ease of reading log
			fmt.Printf("%s: I'm exploring...\n", a.Name)
			time.Sleep(time.Duration(rand.Intn(5)*1000+1000) * time.Millisecond)
			if rand.Float32() < 0.3 {
				fmt.Printf("%s: I have found it!\n", a.Name)
				a.Ship <- "GEM"
				break
			} else {
				fmt.Printf("%s: I will take a break now\n", a.Name)
				a.Ship <- ""
			}
		}
	}
}

// SOLUTION END OMIT
