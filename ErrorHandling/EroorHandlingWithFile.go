package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("iskander.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	bs, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
}
