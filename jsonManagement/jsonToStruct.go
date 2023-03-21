package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	const peopleString string = `[{"firstname":"iskander","lastname":"abbassi","age":24, "is_developer": true },{"firstname":"taher","lastname":"said","age":25,"is_developer": true}]`
	bs := []byte(peopleString)

	people := []Person{}
	err := json.Unmarshal(bs, &people)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", people)
}

type Person struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Age         int8   `json:"age"`
	Information info   `json:"information"`
}

type info struct {
	IsDeveloper bool `json:"is_developer"`
}
