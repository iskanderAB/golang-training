package main

import (
	"fmt"
	"sync"
)
// ensure the result of operation by block access memory with mutex
func main() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		fmt.Println("goroutine 🚀")
		value++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()

	var input string
	fmt.Scanln(&input)

}
