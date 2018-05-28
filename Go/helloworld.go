package main

import "fmt"

func main() {
	var message = "Hello Go World!"
	var greeting *string = &message

	fmt.Println(message, greeting);
	// output is: Hello Go World! 0xc82000a2f0

	fmt.Println(message, *greeting);
	//output is: Hello Go World! Hello Go World!
	
	*greeting = "hi"
	fmt.Println(message, *greeting);
	//output is: hi hi

}