package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs ", runtime.NumCPU())
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var counter int64 = 0
	const gs = 100000
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := counter
			v++
			counter = v
			mutex.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("Go routines ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("counter ", counter)
}
