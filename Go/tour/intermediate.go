package main

import "fmt"

type Rectangle struct {
	x int
	y int
}

func structs() {
	rect := Rectangle{12, 44}
	fmt.Println(rect.x, rect.y)

	another := Rectangle{
		y: 1233,
		x: 2e3,
	}

	fmt.Println(another.x, "rect.x: ", rect.x)
}

func pointers() {
	var1 := 92
	var2 := &var1
	var3 := &var2

	fmt.Println(var3)
	fmt.Println(*var3) // note that this will be the same as the value of var2 which is the address of var1
	fmt.Println(**var3)
	fmt.Println(var2)
	fmt.Println(*var2)
	fmt.Println(var1)
}

func slices() {
	// lets start with arrays in Golang, the first thing you need to know is that they are fixed size and not resizable,
	// which seems very limiting, but don't worry because Go has a way to make them work better

	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a)


	// now lets start with slices

	primes := [6]int{2, 3, 5, 7, 11, 13}
	mySlice := primes[1:4] // this is gonna be [3, 5, 7]

	fmt.Println(mySlice)

	// in the previous example the slicing is happening on the line with the syntax [int:int] where each int is a index of an element of the array to be sliced



	// you can think of slices as references to arrays, meaning that a slice must have an underlying array that it references and once a slice is changed its underlying array is also changed, let me domenstrate:
	names := [4]string{"Moe", "Prachi", "Chris", "Preston"}
	fmt.Println("the original array:", names)

	b := names[0:2]
	c := names[2:4]

	fmt.Println("sliced it up:", b, c)

	b[0] = "NewName"

	fmt.Println("the original array after changing the slice:", names) // NewName is there instead of Moe because the slice mutated its underlying array

	// Note: the slice doesn't actually store any data, instead it actually just just describes a section of the underlying array


	// lets talk about slice literals

	// a slice literal is like an array literal except you don't provide it with the length

        // example:
	// this is an array literal: [3]int{1, 2, 3}
	// and this creates a an array and then builds a slice that references it and returns that: []int{1, 2, 3}

}

// func closures()

func main() {
	// structs()
	// pointers()
	slices()
}

