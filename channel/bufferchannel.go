package main

import "fmt"

func main() {

	c := make(chan int, 2)

	c <- 4
	c <- 10
	fmt.Println(<-c)
	fmt.Println(<-c)
}
