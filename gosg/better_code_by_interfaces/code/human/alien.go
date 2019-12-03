package human

import "fmt"

// START OMIT

type HumanLikeAlien struct {
	Eman string
}

func (a *HumanLikeAlien) SayName() {
	fmt.Println(reverse(a.Eman) + " si eman yM")
}

// END OMIT

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
