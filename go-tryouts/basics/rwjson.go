package basics

import (
	"encoding/json"
	"fmt"
	"log"
)

// use this type to unmarshal json into a struct
type Person struct {
	FirstName string `json:"first_name"`
	LasttName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func RwJson() {
	myJson := `
[
	{
		"first_name": "Clarck",
		"last_name": "Kent",
		"hair_color": "red",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "red",
		"has_dog": false
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println(err)
	}

	log.Printf("unmarshalled %v", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "wally"
	m1.LasttName = "west"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "wally2"
	m2.LasttName = "west2"
	m2.HairColor = "red2"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
