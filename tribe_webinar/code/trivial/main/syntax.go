package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// START NAME OMIT
var count int

type Star struct {
	brightness int
}

func exampleFn() int {
	derivedCount := 0
	return derivedCount
}

// END NAME OMIT

// START FN OMIT
func GreetInVietnamese(seed int64) (greeting string, err error) {
	rand.Seed(seed)
	if rand.Float32() < 0.5 {
		return "", errors.New("I dont know Vietnamese")
	}
	return "Xin Chào!", nil
}

func main() {
	greet, err := GreetInVietnamese(1)
	fmt.Printf(":( %v\n", err)
	fmt.Printf(":) %s\n", greet)
}

// END FN OMIT

// START ARG OMIT

func GetMovies() {}

func GetActionMovies() {}

func GetScifiMovies() {}

func GetHorrorMovies() {}

// END ARG OMIT

// START ARGBOOL OMIT

func GetMovieWithBool(isAction bool, isScifi bool, isHorror bool) {}

// END ARGBOOL OMIT

func ThisIsExported() {}

func thisIsNotExported() {}
