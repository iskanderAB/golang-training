package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sharedState int
	mu          sync.Mutex
	cond        *sync.Cond
)

func main() {

	// Create the condition variable associated with the mutex.
	cond = sync.NewCond(&mu)

	fmt.Println(cond)

	// Start some goroutines to use the condition variable.
	for i := 1; i <= 5; i++ {
		go waitForCondition(i)
	}

	// Simulate some changes to the shared state.
	time.Sleep(time.Second * 5)
	sharedState = 42

	// Signal that the condition has changed.
	cond.Broadcast()

	var input string

	fmt.Scanln(&input)
}

func waitForCondition(id int) {
	cond.L.Lock()
	defer cond.L.Unlock()

	for sharedState != 42 {
		fmt.Printf("Goroutine %d is waiting...\n", id)
		cond.Wait()
	}

	fmt.Printf("Goroutine %d continues!\n", id)
}
