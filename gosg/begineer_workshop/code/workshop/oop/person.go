package oop

import "fmt"

// START OMIT

type Person struct {
	Name string
}

func (p *Person) SayName() {
	fmt.Println("My name is", p.Name)
}

// END OMIT
