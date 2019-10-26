package main

import "fmt"


// here is an example of initializing variables in go
// pretty self explainatory....

var perl, is, awesome = true, true, false;

// note that we ommited the type of the vars,
// go is smart enough to find out the type of the vars

func main() {
	fmt.Println("perl?", perl, "is?", is, "awesome?", awesome)
}
