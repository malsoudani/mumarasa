package main

import (
	"fmt"
	"math/cmplx"
)


// this is a neat trick for wild card declaration too ....
// I know ;) just when you thought you were done with the different ways to do that lol
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// don't worry about the %T and the %v (its just a way of variable string parsing), more on how that works later ...
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}


// go basic types are basically:

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte (alias for uint8)

// rune (alias for int32)
// represents a Unicode code point

// float32 float64

// complex64 complex128
