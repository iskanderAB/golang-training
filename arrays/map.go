package main

import (
	"fmt"
)

func main() {
	var myMap = make(map[string]string)
	var newMap = map[string]int {
		"iskander": 25,
		"ali": 24,
	}

	// test if key exist and run the code 

	if  value, ok := newMap["ali"]; ok {
		fmt.Println("el rojla  => ", value)
	}else{
		fmt.Println("ma fammech rojla :) ")
	}





	myMap["iskander"] = "abbassi"
	myMap["mohammed"] = "khalladi"
	myMap["hammouda"] = "touraa"

	myMap["jj"] = "hahaha"
	
	fmt.Println("map => ", myMap)
	delete(myMap, "iskander")
	fmt.Println("map after delete => ", myMap)
}
