package main

import "github.com/stanleynguyen/slides/dbs-back2school/code/workshop/oop"

// START OMIT

func main() {
	j := &oop.JuniorDeveloper{
		Languages: []string{"go", "ruby", "js"}}
	s := &oop.SeniorDeveloper{}
	internalTraining(j, s, "java", "haskell")
}

func internalTraining(j *oop.JuniorDeveloper, s *oop.SeniorDeveloper, langs ...string) {
	s.Train(j, langs...)
	j.Code()
}

// END OMIT
