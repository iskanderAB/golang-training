package main

import (
	// "math"
	// "reflect"
	"fmt"
	"math/cmplx"
	"time"
)

var (
	uintNumber    uint       = 1 // ma ye9belch (-) nigative number
	intNumber     int64      = 9223372036854775807
	pointerNumber uintptr    = 123154484545454845 //an unsigned integer large enough to store the uninterpreted bits of a pointer value
	floatNumber   float32    = 1.5
	floatNumber64 float64    = 20.5 // yserah barcha
	byteNumber    byte       = 20   // alias for uint8
	complex       complex128 = cmplx.Sqrt(-5 + 12i)
	boolean       bool       = true

// for conversion do int64(value)
)

type Bootcamp struct {
	Lat float64
	// Longitude of the event
	Lon float64
	// Date of the event
	Date time.Time
}

type Point struct {
	X, Y int
}

func main() {

	x := new(Point)
	y := &Point{}

	/* 
	 * declation of struct literals 
	*/
	var (
		p = Point{1, 2}  // has type Point
		q = &Point{1, 2} // has type *Point
		r = Point{X: 1}  // Y:0 is implicit
		s = Point{}      // X:0 and Y:0
	)

	// intNumber  = intNumber + 1
	// fmt.Println("♥ => ",uintNumber);
	fmt.Println(Bootcamp{
		Lat: 34.012836, Lon: -118.495338,
		Date: time.Now(),
	})

	p.X = 5
	fmt.Println("♥ => ", p)
	fmt.Println("♥ => ", q)
	fmt.Println("♥ => ", r)
	fmt.Println("♥ => ", s)
	fmt.Println("☺ => ",  *x , *y)
}