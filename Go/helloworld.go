package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

func CreateMessage(a, b string) string {
	return b + " " + a
}

func Greet(salutation Salutation) { //function definition example
	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

func main() {
	// var message = "Hello Go World!"
	// var greeting *string = &message

	// fmt.Println(message, greeting);
	// output is: Hello Go World! 0xc82000a2f0

	// fmt.Println(message, *greeting);
	//output is: Hello Go World! Hello Go World!
	
	// *greeting = "hi"
	// fmt.Println(message, *greeting);
	//output is: hi hi

	var s = Salutation{"Bob", "wassap"} // example of using user defined type

	Greet(s)


}