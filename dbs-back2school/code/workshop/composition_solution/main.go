package main

import "github.com/stanleynguyen/slides/dbs-back2school/code/workshop/composition"

func main() {
	elonMusk := &composition.AspiredEntrepreneur{
		composition.BusinessGuy{"life on Mars"},
		composition.TechieGuy{"reusable rockets"},
	}
	composition.Pitch(elonMusk)
}
