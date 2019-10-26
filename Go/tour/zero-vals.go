package main

import "fmt"

// variables declared without explicit initial state are given a
// `zero` value .... which is just british for "initial value"
// what are these initial values you say, well
// why don't you run the code to figure out

func main() {
	var i int
	var f float64
	var b bool
	var s string

	// %q is for "quoting"
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
