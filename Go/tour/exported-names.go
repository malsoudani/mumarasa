package main

import (
	"fmt"
	"math"
)

func main() {
	// note that Pi is an "exported name"
	// exported names start with a capital letter
	// when importing a package we can only refer to its exported names

	fmt.Println(math.Pi)

	// if we try to access an unexproted package ....
	// fmt.Println(math.pi)
}
