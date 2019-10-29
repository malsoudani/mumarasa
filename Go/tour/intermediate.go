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
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a);
}

// func closures()

func main() {
	// structs()
	// pointers()
}

