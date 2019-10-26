package main

import (
	"fmt"
	"math"
)

// to "cast" a variable type from something to another thing you use the syntax
// `T(v)` where T refers to the type you want to cast to and v is the value that you
// want casted

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// how would these look like using := construct I wonder?

	fmt.Println(x, y, z)
}


// fun fact:
// Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

// ok after reading it ... it may not be fun for everyone
