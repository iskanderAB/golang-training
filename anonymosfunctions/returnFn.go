package main

import "fmt"

func main() {
	fn := rtnFn()

	fmt.Printf("my result type is %T , value  is %d  ", fn , fn())
}

func rtnFn() func() int {
	return func() int {
		return 100
	}
}
