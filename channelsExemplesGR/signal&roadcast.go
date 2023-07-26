package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Ninja 🥷 explanation

func main() {
	gettingReadyforMession()
}

var ready bool

func gettingReadyforMession() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReady(cond)
	working := 0 

	cond.L.Lock()
	for !ready {
		// time.Sleep(time.Second * 5) // gdima
		working++	
		cond.Wait()
	}
	cond.L.Unlock()

	fmt.Println("number of working ", working)
}

func gettingReady(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano());
	someTime := time.Duration(1 * rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
