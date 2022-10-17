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

	isTrue = false

	if isTrue {
		log.Println("isTrue is, ", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("Cat is ", cat)
	} else {
		log.Println("Cat is not cat.")
	}

	myNum := 100
	// use only one level of this
	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		log.Println("foo")
	} else if myNum == 100 || isTrue {
		log.Println("bar")
	} else if myNum > 1000 && !isTrue {
		log.Println("baz")
	}

	myVar := "Cat"
	// use this if you have more conditions
	switch myVar {
	case "Cat":
		log.Println("Cat is set to cat.")
	case "Dog":
		log.Println("Cat is set to dog.")
	case "Dino":
		log.Println("Cat is set to dino.")
	default:
		log.Println("Nothing to set")
	}

}