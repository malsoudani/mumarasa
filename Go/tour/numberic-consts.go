package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	// which is a really really cool way of writing google or should I say Googol (gimme high five if you got what I mean)
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// Numeric constants are high-precision values.
	// An untyped constant takes the type needed by its context.
	// Try printing needInt(Big) too.
	// (An int can store at maximum a 64-bit integer, and sometimes less.)

	// a numeric constant takes the type of its calling context
	// I know, I know you want to say ... "but you promised me there
	// won't be any explicit casting" and my answer to that ... this
	// is not casting, think of this as pass by value which means that the original
	// type of the numeric constant is preserved just fine try priting them out if you don't believe me
}
