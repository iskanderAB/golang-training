package main

import "fmt"

func main() {
	p1 := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "iskander",
		lastname:  "abbassi",
		age:       25,
	}

	fmt.Println("Anonymos struct => ", p1)
}
