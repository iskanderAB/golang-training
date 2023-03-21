package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var buffer *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Println("please gi me a string ♥ : ")
	input, error := buffer.ReadString('\n')

	if error == nil {
		fmt.Println("you type => ", input)
	}
}
