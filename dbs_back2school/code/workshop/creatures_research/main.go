package main

import (
	"fmt"
	"time"

	"github.com/stanleynguyen/slides/dbs-back2school/code/workshop/creatures_research/research"
	"github.com/stanleynguyen/slides/dbs-back2school/code/workshop/creatures_research/template"
)

func main() {
	creatures := []research.Creature{research.Dog, research.Crab, research.Spider}
	imgSearchInstances := [3]research.ImgSearchInstance{
		research.ImgSearchInstance{},
		research.ImgSearchInstance{},
		research.ImgSearchInstance{},
	}
	eD := research.EyesDetector{}
	lD := research.LegsDetector{}

	start := time.Now()
	template.RunResearchWithoutRoutines(creatures, imgSearchInstances, eD, lD)
	endOfWoRoutines := time.Now()
	fmt.Printf("You took %v without goroutines\n", time.Since(start))
	template.RunResearch(creatures, imgSearchInstances, eD, lD)
	fmt.Printf("You took %v with goroutines\n", time.Since(endOfWoRoutines))
}
