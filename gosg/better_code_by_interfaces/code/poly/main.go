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
	optimusPrime := human.HumanLikeAlien{
		Eman: "Optimus Prime"}
	askForName(&optimusPrime)
}

// COMPARE START OMIT

func askForName(h Human) {
	fmt.Println("What's your name?")
	h.SayName()
}

// PLAY END OMIT

func askForNameExplit(p *human.Person) {
	fmt.Println("What's your name?")
	p.SayName()
}

// COMPARE END OMIT
