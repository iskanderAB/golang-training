package main

import "fmt"
import "strings"

func main() {
	var s string = "hello go"
	var s1 string = " haboub "
	fmt.Println(len(s))
	s1 = s1 + s
	fmt.Println(s1)
	fmt.Println(strings.Compare(s1, s))
	s2 := strings.Replace(s, "h", "O", 1)
	fmt.Println(s2)

	fmt.Println(reverseString("Askander"))

}

func reverseString(str string) string {
	// convert the string to a slice of runes
	runes := []rune(str)
	// reverse the order of the runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// convert the slice of runes back to a string
	fmt.Println(runes)
	return string(runes)
}
