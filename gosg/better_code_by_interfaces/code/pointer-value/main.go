package main

import "fmt"

// DOG START OMIT
type GoodDoggo interface {
	Bark()
	WagTail()
}

type Dog struct {
	Name string
}

func (d Dog) Bark() {
	fmt.Printf("%s: WOOF! WOOF!\n", d.Name)
}

func (d *Dog) WagTail() {
	fmt.Printf("%s: (wags tail) (wags tail)\n", d.Name)
}

// DOG END OMIT

// MAIN START OMIT
func main() {
	myDog := Dog{
		Name: "MiloPeng",
	}
	fmt.Printf("ME: %s, I'm home\n", myDog.Name)
	doggoGreeting(myDog)
}

func doggoGreeting(d GoodDoggo) {
	d.Bark()
	d.WagTail()
}

// MAIN END OMIT
