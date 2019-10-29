package main

import "fmt"

type Rectangle struct {
	x int
	y int
}

func main() {
	rect := Rectangle{12, 44}
	fmt.Println(rect.x, rect.y)

	another := Rectangle{
		y: 1233,
		x: 2e3,
	}

	fmt.Println(another.x)
}
