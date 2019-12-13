package main

import "fmt"

func main() {

	aValue := new(int)

	defer fmt.Println(*aValue) // at the time the reference to aValue is passed in to the deferred call it was still 0 so defer saves this reference and once the rest of the code is done excuting it triggeres with aValue being 0

	for i := 0; i < 100; i++ {
		*aValue++
	}

}
