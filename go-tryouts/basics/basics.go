package basics

import (
	"fmt"
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func Basics() {
	fmt.Println("foo bar")

	var whatToSay string
	var i int

	whatToSay = "foo bar"
	fmt.Println(whatToSay)

	i = 2
	fmt.Println("I is set to", i)

	myMap := make(map[string]string)

	myMap["Dog"] = "Samson"

	log.Println(myMap)
	log.Println(myMap["Dog"])

	otherMap := make(map[string]User)

	me := User{
		FirstName: "Attila",
		LastName:  "Balazs",
	}

	otherMap["me"] = me

	log.Println(otherMap["me"].FirstName)

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

func Decisions() {
	var isTrue bool

	isTrue = true

	if isTrue {
		log.Println("isTrue is, ", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
}
