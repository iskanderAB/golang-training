package main

import "fmt"

func main() {
	// var a [2] string
	// var b = [2]string{"hello", "world!"}
	// c := [...]string{"hello", "world!"}
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// var a [2][3]string
	// for i := 0 ;i < 2; i++ {
	// 	for j := 0 ; j < 3 ; j++ {
	// 		a[i][j] = fmt.Sprintf(" row %d -column %d " , i+1 , j+1)
	// 	}
	// }
	// fmt.Printf("%q" , a)
	// p := []int{2, 3, 4, 5, 6, 13, 14}
	// fmt.Println("all table => ", p)
	// fmt.Println(" slice table =>  ", p[0:3])
	// fmt.Println(" left slice => ", p[:3])
	// fmt.Println(" right slice ", p[4:])

	// println("--------------------")

	// cities := make([]string, 3)
	// cities[0] = "bizerte"
	// cities[1] = "beja"
	// cities[2] = "tatwin"
	// fmt.Println("cities => ", cities)
	// fmt.Printf("%q \n", cities)
	// streets := []string{}
	// streets = append(streets, "hasn el Nouri")
	// streets = append(streets, "madina monawra ")
	// fmt.Println("streets =>" , streets)

	// println("--------------------")

	// cities := []string{"San Diego", "Mountain View"}
	// otherCities := []string{"Santa Monica", "Venice"}
	// cities = append(cities, otherCities...)
	// fmt.Printf("all cities => %q \n", cities)
	// fmt.Println("length of array is => ", len(cities))

	// println("--------------------")

	// var z []int
	// fmt.Println( z , len(cities) , cap(cities))

	// if z == nil {
	// 	fmt.Println("nil !")
	// }

	// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// for i, v := range pow { //index , value
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	// a := 10 // 1010
	// b := 3  // 0011

	// fmt.Println(" a and b binary => ", a&b) // 1010 and 0011 = 0010 => decimal 2
	// fmt.Println(" a or b binary => ", a|b)  // 1010 or 0011 = 1011 => decimal 11
	// fmt.Println(" a xor  ", a^b)           // 1010 xor 0011 = 1001 => decimal 9
	// fmt.Println(" a xand  ", a&^b)         // 1010 xand 0011 = 0100=> decimal 8

	// Bitwise Operators, Bit Shifting

	// var byteNumber byte = 8

	// byteNumber = byteNumber << 4 //  shift 3 bit "8 => 2*3" * 2*4 => 2*7 =  128
	// byteNumber = byteNumber >> 3 // 128 => 2*7 / 2*3 => 2*4 = 16

	// fmt.Println(" << operations >> ==> ", byteNumber)

	// pow := make([]int, 10)
	// for i := range pow {
	// 	pow[i] = 1 << uint(i)
	// }
	// for _, value := range pow {
	// 	fmt.Printf("%d\n", value)
	// }

	cities := map[string]int{
		"bizerte" : 833374,
		"souuse" : 4578945,
		"sfax" : 749656,
	}

	for key,value := range cities {
		fmt.Printf("%s has  %d inhabitants" , key , value)
	}







}
