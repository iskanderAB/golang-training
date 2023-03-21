package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(" ******* beginning *********")
	wg.Add(16)

	go iskanderGR()
	go samiGR()
	go zahmoulaGR()

	fmt.Println(" ******* statistics *********")
	fmt.Println("OS : \t\t", runtime.GOOS)
	fmt.Println("Arch : \t\t", runtime.GOARCH)
	fmt.Println("CPUs : \t\t", runtime.NumCPU())
	fmt.Println("GoRoutines  : \t", runtime.NumGoroutine())
	wg.Wait()
}

func iskanderGR() {
	for i := 0; i < 1000000000; i++ {
		fmt.Println("iskander go routine => \t\t", i)
		time.Sleep(time.Second)
	}
	wg.Done()
}
func samiGR() {
	for i := 0; i < 10; i++ {
		fmt.Println("sami go routine => \t\t", i)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func zahmoulaGR() {
	for i := 0; i < 10; i++ {
		fmt.Println("zahmoula go routine => \t\t", i)
		time.Sleep(time.Second)
	}
	wg.Done()
}
