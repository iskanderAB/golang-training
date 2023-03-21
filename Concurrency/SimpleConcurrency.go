package main

import (
	"bufio"
	"fmt"
	"os"
)

var i int = 0

// simple explanation
func main() {
	for j := 0; j < 100000000000; j++ {
		ConcurrencyApplication()
	}

	var buffer *bufio.Reader = bufio.NewReader(os.Stdin)
	input, error := buffer.ReadString('\n')

	if error == nil {
		fmt.Println("you type => ", input)
	}

}
func ConcurrencyApplication() {
	fmt.Println(" increment => ", i)
	i++
}
