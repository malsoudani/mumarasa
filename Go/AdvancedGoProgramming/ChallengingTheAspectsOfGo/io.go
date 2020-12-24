package main

import (
	"log"
)

func main() {
	thing := 0
	log.Println("log.Prinln: ")
	for i := 0; i < 100; i += 1 {
		thing += 1
		go log.Println(thing)
	}
	// fmt.Println("fmt.Prinln: ")
	// for i := 0; i < 100; i += 1 {
	// 	go fmt.Println(i)
	// }
}
