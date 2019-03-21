package models

// Xall
type XAll struct {
	Persons []Person
	Animals []Animal
}

// Person
type Person struct {
	Name   string
	Mobile string
}

// Animal
type Animal struct {
	AnimalID string
	Animal   string
}
