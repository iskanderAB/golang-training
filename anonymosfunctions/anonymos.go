package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello anonymos :p ");
	}()

	f := func ()  {
		fmt.Println("hello from exp fn ");
	} 

	f()
}