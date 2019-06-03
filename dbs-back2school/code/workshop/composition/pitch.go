package composition

import "fmt"

// START OMIT

func Pitch(f Founder) {
	fmt.Println("INVESTORS: What do you have to offer?")
	fmt.Printf("FOUNDER: I'm going to achieve %s\nwith %s\n", f.BizIdea(), f.TechIdea())
}

// END OMIT
