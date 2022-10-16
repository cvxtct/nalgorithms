package basics

import (
	"fmt"
	"log"
)

func Basics() {
	fmt.Println("foo bar")

	var whatToSay string
	var i int

	whatToSay = "foo bar"
	fmt.Println(whatToSay)

	i = 2
	fmt.Println("I is set to", i)
}

func PointerBasics() {

	// var myString string
	// myString = "Green"
	myString := "Green"

	log.Println("before:", myString)
	changeUsingPointer(&myString)
	log.Println("after:", myString)

}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
