package main

import "fmt"

func main() {

	var c = make(chan string)

	go send(c)
	go recive(c)

	var input string
	fmt.Scanln(&input)
}

func send(c chan<- string) {
	c <- "hello iskander ♥ "
}

func recive(c <-chan string) {
	message := <-c
	fmt.Println("message from sender ", message)
}
