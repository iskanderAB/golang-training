package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mouin dog = 10
	var numberString = "100"
	fmt.Printf("hello %v ", mouin)
	newNumber, _ := strconv.ParseInt(numberString, 10, 8)
	fmt.Println(newNumber)
}

type dog int
