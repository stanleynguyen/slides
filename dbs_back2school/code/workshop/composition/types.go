package composition

// START OMIT

type BusinessGuy struct {
	BusinessIdea string
}

func (g *BusinessGuy) BizIdea() string {
	return g.BusinessIdea
}

type TechieGuy struct {
	SpecialTech string
}

func (g *TechieGuy) TechIdea() string {
	return g.SpecialTech
}

type AspiredEntrepreneur struct {
	BusinessGuy
	TechieGuy
}

// END OMIT
