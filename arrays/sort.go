package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var intsSlice = []int{14, 12, 3, 4, 0, 2}
	fmt.Fprintln(os.Stdout, intsSlice)
	sort.Ints(intsSlice)
	fmt.Fprintln(os.Stdout, intsSlice)

}
