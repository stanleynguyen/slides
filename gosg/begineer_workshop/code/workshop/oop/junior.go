package oop

import "fmt"

// START OMIT

type JuniorDeveloper struct {
	Languages []string
}

func (d *JuniorDeveloper) Code() {
	langs := ""
	for _, v := range d.Languages {
		langs += v + ", "
	}
	langs = langs[:len(langs)-2]
	fmt.Println("I code in " + langs)
}

// END OMIT
