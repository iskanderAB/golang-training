package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := Person{
		Firstname: "iskander",
		Lastname:  "abbassi",
		Age:       24,
	}
	p2 := Person{
		Firstname: "taher",
		Lastname:  "said",
		Age:       25,
	}

	people := []Person{p1, p2}
	data, err := json.MarshalIndent(people, "", "  ")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("data => ", string(data))
}

type Person struct {
	Firstname string
	Lastname  string
	Age       int8
}
