package main

import "fmt"

func main() {

	person := User{"iskander", 23,true}

	// person.age = 20
	// person.name = "isakander"
	// person.status = true

	fmt.Println(person)
}

type User struct {
	name   string
	age    int8
	status bool
}
