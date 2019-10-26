package main

import "fmt"


// a paradigm that Golang has followed with typing their variables/function is
// placing the type after variable/function to refer to their type

// notice that x is of type int, and so is y, .... and so is add()


func add(x int, y int) int {
	return x + y
}

// an alternate way to declare multiple variables at once is ommiting the type from all of the variables but the last:
// we can do something like func add(x, y int) int ....

func main() {
	fmt.Println(add(42, 343)) // 385
}

