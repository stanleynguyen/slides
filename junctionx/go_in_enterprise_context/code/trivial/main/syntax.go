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
	if err != nil {
		fmt.Printf(":( %s\n", err.Error())
		return
	}
	fmt.Printf(":) %s\n", greet)
}

// END FN OMIT

func ThisIsExported() {}

func thisIsNotExported() {}
