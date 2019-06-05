package research

import (
	"fmt"
	"math/rand"
	"time"
)

// Creature types for creatures in nature
type Creature uint8

const (
	// Dog a dog
	Dog = Creature(iota)
	// Crab a crab
	Crab
	// Spider spidey
	Spider
)

const photosFromGoogle = " photos from Google"

func sleep() {
	sleep := time.Duration(rand.Intn(10) * 100)
	time.Sleep(sleep * time.Millisecond)
}

// ImgSearchInstance Google instace for image searching
type ImgSearchInstance struct{}

// Search search for creatures images
func (i *ImgSearchInstance) Search(c Creature) string {
	sleep()
	switch {
	case c == Dog:
		return string(Dog) + photosFromGoogle
	case c == Crab:
		return string(Crab) + photosFromGoogle
	case c == Spider:
		return string(Spider) + photosFromGoogle
	}
	return "UNKNOWN"
}

// EyesDetector deep learning instace to detect creatures' eyess
type EyesDetector struct{}

// Detect detect a creature's eyes
func (d *EyesDetector) Detect(googleRs string) int {
	sleep()
	switch {
	case googleRs == string(Dog)+photosFromGoogle:
		fallthrough
	case googleRs == string(Crab)+photosFromGoogle:
		return 2
	case googleRs == string(Spider)+photosFromGoogle:
		return 8
	}
	return 0
}

// LegsDetector deep learning instace to detect creatures' legs
type LegsDetector struct{}

// Detect detect a creature's legs
func (d *LegsDetector) Detect(googleRs string) int {
	sleep()
	switch {
	case googleRs == string(Dog)+photosFromGoogle:
		return 4
	case googleRs == string(Crab)+photosFromGoogle:
		fallthrough
	case googleRs == string(Spider)+photosFromGoogle:
		return 8
	}
	return 0
}

// PrintResult Announce research results for a creature
func PrintResult(c Creature, eyes, legs int) {
	var name string
	switch {
	case c == Dog:
		name = "Dogs"
	case c == Crab:
		name = "Crabs"
	case c == Spider:
		name = "Spider"
	}
	fmt.Printf("%s have %v eyes and %v legs\n", name, eyes, legs)
}
