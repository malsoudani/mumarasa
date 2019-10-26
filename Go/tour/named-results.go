package main

import "fmt"

// with multiple returns comes the feature of named multiple returns
// which allows your function to do a "naked" return and it will automagically
// return the last value assigned to the named return values in the context of
// the called function

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := split(30)
	fmt.Println(a, b)
}
