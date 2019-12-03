package main

import (
	"fmt"

	"github.com/stanleynguyen/slides/gosg/better_code_by_interfaces/code/human"
)

// START OMIT

type Dog struct{}

func (d *Dog) Bark() {
	fmt.Println("WOOF")
}

func main() {
	pete := human.Person{
		Name: "Peter"}
	askForName(&pete)
	dog := Dog{}
	askForName(&dog)
}

func askForName(h interface{}) {
	fmt.Println("What's your name?")
	p := h.(*human.Person)
	p.SayName()
}

// END OMIT
