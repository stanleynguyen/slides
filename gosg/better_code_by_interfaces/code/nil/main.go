package main

import (
	"fmt"
	"os"
)

// START OMIT
func TypedError() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := TypedError()
	fmt.Printf("is eq to nil: %v\n", err == nil)
	fmt.Printf("is eq to (*os.PathError)(nil): %v\n", err == (*os.PathError)(nil))
}

// END OMIT
