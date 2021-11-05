// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		name , age , locatioin  = "iskander" , 23 , "el alia" 
	)
	// const (
	// 	pi = 3.14
	// )
	const pi = 3.14 // you cant create  const with := and const can't be a func :'(	
	/*Constants can only be character, string, boolean, or numeric values and
	cannot be declared using the := syntax */

	const (
		big =  1 << 2
 	)



	action := func ()  {
		 fmt.Print("Hi iskander \n");
	}
	action();

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(locatioin))
	
	fmt.Print(big)
}
