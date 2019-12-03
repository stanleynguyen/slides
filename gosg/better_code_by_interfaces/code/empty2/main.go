package main

import (
	"fmt"

	"github.com/stanleynguyen/slides/gosg/better_code_by_interfaces/code/human"
)

// START OMIT

type Dog struct{}

func (d *Dog) Bark() { fmt.Println("WOOF") }

func main() {
	pete := human.Person{Name: "Peter"}
	askForName(&pete)
	dog := Dog{}
	askForName(&dog)
	askForName(100)
}

func askForName(h interface{}) {
	fmt.Println("What's your name?")
	if p, ok := h.(*human.Person); ok {
		p.SayName()
	} else if d, ok := h.(*Dog); ok {
		d.Bark()
	} else {
		fmt.Println("I dont know")
	}
}

// END OMIT
