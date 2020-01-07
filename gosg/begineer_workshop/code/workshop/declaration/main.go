package main

import "fmt"

func main() {
	//START OMIT
	var x int     // Variable x initialized with zero-value
	var y int = 2 // y with value 2
	z := 3        // implicit type declaration
	y, z = x, y   // double declaration

	var b bool
	var st string

	s := make([]int, 3) // slice declaration with capacity 3
	fmt.Println(x, b, st, s)
	// END OMIT

	y = z
}
