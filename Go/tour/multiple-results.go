package main

import "fmt"


// a function in go can return any number of results,
// the swap function utilizes this idea here:

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("foo", "bar")
	fmt.Println(a, b)
}
