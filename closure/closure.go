package main

import "fmt"

func main() {
	increment := closureFunction()
	increment1 := closureFunction()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(increment1())
	fmt.Println(increment1())
	fmt.Println(increment1())
}

func closureFunction() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
