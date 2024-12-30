package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("We are learning JSON")
	person := Person{Name: "Aakash", Age: 21, IsAdult: true}
	fmt.Println(person)

	//convert Person into JSON encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("person data is", string(jsonData))

	//Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarsahlling", err)
		return
	}
	fmt.Println("person data is", personData)
}
