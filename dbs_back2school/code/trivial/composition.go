package trivial

// START OMIT

type ExampleInterface interface {
	Read(b []byte) (string, error)
	Write(dest string) error
}

type AnotherExampleInterface interface {
	Echo() string
}

// END OMIT
