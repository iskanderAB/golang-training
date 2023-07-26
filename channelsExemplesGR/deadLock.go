package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	accessMemory    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
	
		v1.accessMemory.Lock()
		fmt.Println("duration go ...")
		time.Sleep(2 * time.Second)
		v2.accessMemory.Lock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)

		wg.Done()
		v1.accessMemory.Unlock()
		v2.accessMemory.Unlock()
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()

}
