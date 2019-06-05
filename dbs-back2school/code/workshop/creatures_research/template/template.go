package template

import "github.com/stanleynguyen/slides/dbs-back2school/code/workshop/creatures_research/research"

// RunResearch carry out creatures research with goroutines
func RunResearch(
	creatures []research.Creature,
	imgInstances [3]research.ImgSearchInstance,
	eD research.EyesDetector,
	lD research.LegsDetector) {

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
