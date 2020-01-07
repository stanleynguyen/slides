package main

import (
	"fmt"
	"reflect"
)

// START OMIT
type basicTypes struct {
	b   bool
	i   int
	i64 int64
	ui  uint64
	f   float32
	s   string
}

func main() {
	t := basicTypes{true, 1, 1e12, 1234, 1.234, "I am a string"}
	printTypeStruct(t)
}

// END OMIT

func printTypeStruct(t basicTypes) {
	printType(t.b)
	printType(t.i)
	printType(t.i64)
	printType(t.ui)
	printType(t.f)
	printType(t.s)
}

func getTypeString(t interface{}) reflect.Kind {
	return reflect.ValueOf(t).Kind()
}

func printType(v interface{}) {
	fmt.Printf("%v is of type %s\n", v, getTypeString(v))
}
