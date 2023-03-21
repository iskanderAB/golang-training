package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := make(chan string)

	go func() {
		for i := 1; i <= 100; i++ {
			c <- "hello sknder ☻ " + strconv.Itoa(i)
		}
		close(c) // important for avoid deadlock goroutine
	}()

	for v := range c {
		fmt.Println("value => ", v)
	}
}
