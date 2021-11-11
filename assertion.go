package main

import (
	"fmt"
	"time"
)

func timeMap(y interface{}) {
	z, ok  := y.(map[string]interface{})
	println("debugger =>  z ", ok)
	if ok {
		z["updated_at"] = time.Now()
		z["iskander"] = "abbassi"
	}
}

func main() {
	foo := map[string]interface{}{
		"Matt": 42,
	}
	timeMap(foo)
	fmt.Println(foo)
}
