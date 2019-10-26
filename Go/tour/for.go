// lets talk looping in go

// Go has only one looping construct, the for loop.

// The basic for loop has three components separated by semicolons:

// the init statement: executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement: executed at the end of every iteration
// The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

// The loop will stop iterating once the boolean condition evaluates to false.

// Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}


// the init and the post statements are actually optional so you can do something like:
// for ; sum < 1000; {
// 	sum += sum
// }

// well at this point you can just drop the semicolons because C's `while` is basically spelled `for` in Go
// for sum < 1000 {
// 	sum += sum
// }

// remember that trick your cs teacher showed you to exceed the output limit or time limit with an infinite while loop:
// well here is how you do this with Go:
// for {
// }
