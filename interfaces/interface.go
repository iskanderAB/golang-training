package main

import "fmt"

func main() {
	fmt.Println("begin")
	iskander := person{
		name: "iskander",
		age: 24,
		weight: 82,
	}
	
	var taher human 

	fmt.Println("Selected person => ", iskander)

	iskander.speack("sami widir")

}

type person struct {
	name   string
	age    int
	weight float32
}

func (p person) speack(word string) {
	fmt.Println("hello ", word)
}

type human interface {
	speack()
}