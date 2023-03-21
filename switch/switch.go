package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)


	input, _ := reader.ReadString('\n');
	
	number , _ := strconv.ParseFloat(strings.TrimSpace(input),32);


	switch number {
	case 1:
		fmt.Println("it is 1 !")
	case 2:
		fmt.Println("it is 2 !")
		fallthrough
	case 3:
		fmt.Println("it is 3 !")
		fallthrough
	case 4:
		fmt.Println("it is 4 !")
	case 5:
		fmt.Println("it is 5 !")
	default :
		fmt.Println("not a number :) ")
	}

}