package main

import (
	"fmt"
	"math"
)

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("this function return => ", geString())
	arg, integer := getDoubleValue("hello iskander")
	fmt.Println("double return => ", arg, integer)
	variadicFunction("isdkander", "hammdi", "hamdi", "sami")
	unfurlingSlice(slice...)
	fmt.Printf("%T , %v ", returnFunction(), returnFunction()())
	fmt.Println(functionWithCallbacks(func(x float64) bool {
		return math.Mod(x, 2) == 0
	}, 10))
}

func geString() string {
	return fmt.Sprint("hello world")
}

func getDoubleValue(arg string) (string, int) {
	return arg, 20
}

func variadicFunction(x ...string) {
	fmt.Println("all args => ", x)
}

func unfurlingSlice(args ...int) {
	fmt.Println("args => ", args)
}
func returnFunction() func() string {
	return func() string {
		return fmt.Sprint(" returned from sub function ")
	}
}

func functionWithCallbacks(isSomething func(x float64) bool, x float64) string {
	if isSomething(x) {
		return fmt.Sprintf("\n opertation with %f is true ", x)
	} else {
		return fmt.Sprintf("\n opertation with %f is false ", x)
	}
}
