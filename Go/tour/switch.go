// Switch statement in go

// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow.
// In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go.
// Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

package main

import (
	"fmt"
	"runtime"
)


func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

// example:

// t := time.Now()
// switch {
// case t.Hour() < 12:
// 	fmt.Println("Good morning!")
// case t.Hour() < 17:
// 	fmt.Println("Good afternoon.")
// default:
// 	fmt.Println("Good evening.")

// Notes:
// 	a couple of notes to take from the previous example is that go can have expressions as its cases as you can see
// 	and also if a case is met then it would not go on to check for the next case, which is in effect means that the `break`
// 	statement needed in other programming languages is done automatically in go.
// 	also another thing to note is that a switch with no condition would be the equavilant to `switch true { .... }` meaning that
// 	this construct is a great replacement to long chained if else statement used in other programming language without having to
// 	plan for using it.
