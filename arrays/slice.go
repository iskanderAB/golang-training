package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	var slice = []string{"iskander", "ali", "khalladi", "mohammed", "mouin", "sami", "dali"}
	slice = append(slice[:2])
	fmt.Printf("type is %T , value is %s \n", slice, slice)

	//sort function
	var intSlice = []int{1, -2, 4, 61}

	// sort.Ints(intSlice);
	fmt.Println("mu nuber slice is ", intSlice)
	fmt.Println("is sorted ? :  ", sort.IntsAreSorted(intSlice))
	executionTime := time.Since(start)

	var broSlice = []string{"iskander", "ali", "khalladi", "mohammed", "mouin", "sami", "dali"}

    broSlice = append(broSlice[:2],broSlice[3:]...)

	fmt.Println("slice after cut => ", broSlice);
	fmt.Println("execution time is equal => ", executionTime);
}
