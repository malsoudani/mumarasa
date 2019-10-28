// if statements in golang

// if statement, much like the for loop, the condition doesn't need to be surrounded by parens but the body requires curly braces.

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}


// the if statement, much like the for loop can start with a short statement before the condition, something like this:
// if v := math.Pow(x, n); v < lim {
// 	return v
// }
