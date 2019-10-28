// what is the defer statement in go?
// the defer statement is a way to excute a function after the surrounding function return
// the deferred call's arguments are evaluated immediately but the function excution with these arguments is held off until the surrounding function finishes its job.

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}


// when defer satements get called multiple times in a single function they actually get pushed into a stack and get excuted in a LIFO (Last In First Out) order

// comment out the top example and uncomment this one

// package main

// import "fmt"

// func main() {
// 	fmt.Println("counting")

// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}

// 	fmt.Println("done")
// }
