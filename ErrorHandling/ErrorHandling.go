package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no error => ", n)

	file, err := os.Create("iskander.txt")

	if err != nil {
		fmt.Println("error => ", err)
		return
	}
	defer file.Close()

	r := strings.NewReader("hello iskander from CJS")
	io.Copy(file, r)
}
