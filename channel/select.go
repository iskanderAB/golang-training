package main

import (
	"fmt"
	"strconv"
)

func main() {
	even := make(chan string)
	odd := make(chan string)
	quit := make(chan string)

	go send(even, odd, quit)
	go recive(even, odd, quit)

	var input string
	fmt.Scanln(&input)

}

func send(even, odd, quit chan<- string) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- "even ☺ " + strconv.Itoa(i)
		} else {
			odd <- "odd ☻ " + strconv.Itoa(i)
		}
	}
	quit <- "♪ clawsed ♪"
}

func recive(even, odd, quit <-chan string) {
	for {
		select {
		case v := <-even:
			fmt.Println(v)
			break
		case v := <-odd:
			fmt.Println(v)
			break
		case v := <-quit:
			fmt.Println(v)
			return
		}
	}
}
