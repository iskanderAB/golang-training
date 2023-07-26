package main

import (
	"fmt"
	"time"
)

func main() {
	tk := time.NewTicker(time.Second)


	i := 0
	for range tk.C { 
		i++
		fmt.Println("tick ⚡")
		if i > 3 { tk.Stop(); break}
	}
}