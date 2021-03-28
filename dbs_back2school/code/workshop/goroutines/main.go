package main

import (
	"fmt"
	"time"
)

// START OMIT
func deplayedPrint(i int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(i)
}

func main() {
	start := time.Now()
	for i := 1; i < 6; i++ {
		deplayedPrint(i)
	}
	fmt.Printf("Program took %s\n", time.Since(start))
}

// END OMIT
