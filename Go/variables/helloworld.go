package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

func CreateMessage(a, b string) (string, string) {
	return b + " " + a, "Hey " + a
}

func Greet(salutation Salutation, do Printer) { //function definition example
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	do(message)
	do(alternate)
}

type Printer func(string) () 	//declaring a custom function type as with a specific blueprint
								//pretty much here we are saying that Printer redefines the blueprint func(string) ()
func Print(s string) {
	fmt.Print(s)
}

func Println(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {	// creating a function that returns a Printer
	return func(s string) {							// which is a closure defined within the parent
		fmt.Println(s + custom)						// that uses the resources of the parent function
	}												// and returns the value of that function
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

	Greet(s, CreatePrintFunction("!!!!"))

}