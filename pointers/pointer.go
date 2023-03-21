package main

import "fmt"

func main() {

	var pointer *int

	var number int = 20

	pointer = &number

	*pointer = 10

	fmt.Println("my number is equal => ", number)

}
