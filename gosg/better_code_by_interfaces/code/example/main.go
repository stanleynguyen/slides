package main

import (
	"fmt"

	"github.com/stanleynguyen/slides/gosg/better_code_by_interfaces/code/human"
)

// EXAMPLE START OMIT

type Human interface {
	SayName()
}

// EXAMPLE END OMIT

// PLAY START OMIT

func main() {
	pete := human.Person{
		Name: "Peter"}
	askForName(&pete)
}

func askForName(h Human) {
	fmt.Println("What's your name?")
	h.SayName()
}

// PLAY END OMIT
