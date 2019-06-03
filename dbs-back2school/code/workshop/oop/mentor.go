package oop

// START OMIT

type SoftwareMentor struct{}

func (m *SoftwareMentor) Train(d *JuniorDeveloper, langs ...string) {
	for _, l := range langs {
		d.Languages = append(d.Languages, l)
	}
}

// END OMIT
