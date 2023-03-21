package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 10)
		c <- 40
	}()
	fmt.Println("ready to jump ")
	go func() {
		fmt.Println(" hello channel => ", <-c)
	}()
	fmt.Println(" jumped ")

	var input string
	fmt.Scanln(&input)
}
