package solution

import "github.com/stanleynguyen/slides/dbs-back2school/code/workshop/creatures_research/research"

// RunResearch carry out creatures research with goroutines
func RunResearch(
	creatures []research.Creature,
	imgInstances [3]research.ImgSearchInstance,
	eD research.EyesDetector,
	lD research.LegsDetector) {

	eyesRecord := map[research.Creature]chan int{}
	legsRecord := map[research.Creature]chan int{}
	for _, c := range creatures {
		eyesRecord[c] = make(chan int)
		legsRecord[c] = make(chan int)
		go launchImgSearchInstances(imgInstances, eD, lD, c, eyesRecord[c], legsRecord[c])
	}
	for i := 0; i < len(creatures); i++ {
		c := creatures[i]
		research.PrintResult(c, <-eyesRecord[c], <-legsRecord[c])
	}

}

// RunResearchWithoutRoutines carry out creatures research without goroutines
func RunResearchWithoutRoutines(
	creatures []research.Creature,
	imgInstances [3]research.ImgSearchInstance,
	eD research.EyesDetector,
	lD research.LegsDetector) {

	for i := 0; i < len(creatures); i++ {
		imgRs := imgInstances[0].Search(creatures[i])
		eyeRs := eD.Detect(imgRs)
		legRs := lD.Detect(imgRs)
		research.PrintResult(creatures[i], eyeRs, legRs)
	}

}

func launchImgSearchInstances(
	imgInstances [3]research.ImgSearchInstance,
	eD research.EyesDetector,
	lD research.LegsDetector,
	c research.Creature,
	eCh, lCh chan int) {

	imgRsChs := [3]chan string{
		make(chan string),
		make(chan string),
		make(chan string),
	}
	for i, instance := range imgInstances {
		go func(rsCh chan string, instance research.ImgSearchInstance) {
			rsCh <- instance.Search(c)
		}(imgRsChs[i], instance)
	}
	v := ""
	for v == "" {
		select {
		case v = <-imgRsChs[0]:
			saveEyesLegsRecord(c, v, eD, lD, eCh, lCh)
		case v = <-imgRsChs[1]:
			saveEyesLegsRecord(c, v, eD, lD, eCh, lCh)
		case v = <-imgRsChs[2]:
			saveEyesLegsRecord(c, v, eD, lD, eCh, lCh)
		default:
			continue
		}
	}

}

func saveEyesLegsRecord(
	c research.Creature,
	imgRs string,
	eD research.EyesDetector,
	lD research.LegsDetector,
	eCh, lCh chan int) {

	eyeRs := make(chan int)
	legRs := make(chan int)
	go func(rsCh chan int) {
		rsCh <- eD.Detect(imgRs)
	}(eyeRs)
	go func(rsCh chan int) {
		rsCh <- lD.Detect(imgRs)
	}(legRs)
	eye := <-eyeRs
	leg := <-legRs
	eCh <- eye
	lCh <- leg

}
