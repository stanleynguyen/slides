package main

import "github.com/stanleynguyen/slides/dbs-back2school/code/workshop/oop"

// EXAMPLE START OMIT

type Developer struct {
	UnderstandingOfTechnology float64
}

type CommunityBuilder struct{}

func (b *CommunityBuilder) BuildCommunity() {}

type DeveloperAdvocate struct {
	Developer        // DeveloperAdvocate.UnderstandingOfTechnology is valid
	CommunityBuilder // DeveloperAdvocate.BuildCommunity() is a valid call
}

// EXAMPLE END OMIT

// MAIN START OMIT

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

// MAIN END OMIT
