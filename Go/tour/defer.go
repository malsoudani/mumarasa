// what is the defer statement in go?
// the defer statement is a way to excute a function after the surrounding function return
// the deferred call's arguments are evaluated immediately but the function excution with these arguments is held off until the surrounding function finishes its job.

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
