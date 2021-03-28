package main

import "fmt"

// START FOUNDER OMIT

type Founder interface {
	BizIdea() string
	TechIdea() string
}

// END FOUNDER OMIT

// START PITCH OMIT

func Pitch(f Founder) {
	fmt.Println("INVESTORS: What do you have to offer?")
	fmt.Printf("FOUNDER: I'm going to achieve %s\nwith %s\n", f.BizIdea(), f.TechIdea())
}

// END PITCH OMIT

// START TYPE OMIT

type BusinessGuy struct {
	BusinessIdea string
}

func (g *BusinessGuy) BizIdea() string {
	return g.BusinessIdea
}

type TechieGuy struct {
	SpecialTech string
}

func (g *TechieGuy) TechIdea() string {
	return g.SpecialTech
}

type AspiredEntrepreneur struct {
	BusinessGuy
	TechieGuy
}

// END TYPE OMIT

// START MAIN OMIT
func main() {
	elonMusk := AspiredEntrepreneur{
		BusinessGuy{"life on Mars"},
		TechieGuy{"reusable rockets"},
	}
	Pitch(&elonMusk)
}

// END MAIN OMIT
