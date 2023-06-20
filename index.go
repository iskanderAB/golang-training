package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error => ", err.Error())
	}
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Println("resp => ", string(content))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("clawzed :p")
		}
	}(resp.Body)
}
