package main

// importing packages in go, Note: this is called a factored import statement while we can still import each of these packages individually
import (
	"fmt"
	"math/rand"
)

func main() {
	// environemnt deterministic programs meaning that everytime I run rand on the same environment I will get the same thing over and over
	fmt.Println("My Fav number is", rand.Intn(10))

	// we need to seed to get a initial deterministic state of our own
	rand.Seed(4)
	fmt.Println("My Own Fav Number is", rand.Intn(10))
}
