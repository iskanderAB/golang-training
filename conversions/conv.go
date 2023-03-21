package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader bufio.Reader = *bufio.NewReader(os.Stdin)

	fmt.Print("please give ma a number: ")
	input, _ := reader.ReadString('\n')

	number, err := strconv.ParseFloat(strings.TrimSpace(input), 32)

	if err == nil {
		fmt.Println("input is ", number+1)
	} else {
		panic(err)
	}
}
