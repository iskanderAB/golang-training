package main

import "fmt"

func main() {
	str1 := secretAgent{
		person: person{
			lastname:  "abbassi",
			firstname: "iskander",
			age:       24,
		},
		ltk: true,
	}

	str1.speak()
}

type person struct {
	firstname string
	lastname  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Printf("hello form speaker I'm %s %s I have %v years old , my status is %v", s.firstname, s.lastname, s.age, s.ltk)
}

type human interface{ 
	speak()
}
