package main

import "fmt"

// short varianle declarations:
// with all of the ways that we learned how to declare variables in go
// the simplist and shortest has to be using the `:=` construct,
// which is called an implict declarations because it is considered implied
// that you want a variable with what ever the type of the thing that you place
// on the right hand side ... well lets try it:

// perl := "is awesome"

// aaaand if you uncommented the last line and ran the code it would break
// because everything outside of the scope of a function is expected to be
// prefixed with something that tells go how to process it, so the parser
// looks for something prefixed with var, func and so on.

// so go is pretty smart but not smart enough for this ...

// so you have to stick with declaring this way outside of the scope of a function:

var perl = "its awesome"

func main() {
	fmt.Println("perl?", perl)

	// however we can use the `:=` construct just fine in a function scope:

	golang := "is awesomer" // not using go here because its a reserved keyword .. more on that later
	fmt.Println("go?", golang)
}
