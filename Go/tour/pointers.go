package main

import "fmt"

// pointers in Go are very similar to these in c++ and c:
// you access the address of a varibale in the memory by the &myVariable syntax
// you dereference a variable through its address and access its content using the *myAddress syntax

func main() {
	i, j := 34, 56
	p, k := &j, &i
	fmt.Println(p);
	fmt.Println(*k);

	// setting throught the pointer

	*p = 12

	fmt.Println(j);

	*k = *p
	fmt.Println(i);
}

// one note: unlike C, Go has no pointer arithmetic
