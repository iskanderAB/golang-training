package main


import "fmt"

func add(x int, y int) int {
	return x + y
}

func multReturn(x int, y int) (int, int)  { // u can return multiple result ♥ 
	return x , y
}


func location(city string) (a, b string) { // trajja3 direct a wel b elli declarithom fel params 
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		a, b = "California", "North America"
	case "New York", "NYC":
		a, b = "New York", "North America"
	default:
		a, b = "Unknown", "Unknown"
	}
	
	return
}

func main() {
	println(add(5,6))
	println(multReturn(2,5))
	region1, continent := location("Santa Monica")
	fmt.Printf("Matt lives in %s, %s", region1, continent)
}