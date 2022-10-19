package basics

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func InterFacesPractice() {
	dog := Dog{
		Name:  "doggo",
		Breed: "German Shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "gro",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)

}

func PrintInfo(a Animal) {
	fmt.Println("this animal says", a.Says(), " and has ", a.NumberOfLegs())
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Woof"
}

func (d *Gorilla) NumberOfLegs() int {
	return 4
}
