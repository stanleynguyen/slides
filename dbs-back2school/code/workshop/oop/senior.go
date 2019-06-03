package oop

import "fmt"

// START OMIT
type SeniorDeveloper struct {
	SoftwareMentor
}

func (d *SeniorDeveloper) Code() {
	fmt.Println("I dont only code. I solve problems")
}

// END OMIT
