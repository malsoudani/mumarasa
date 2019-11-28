package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	log.Println("log Entry")
	log.SetFlags(0)

	for i := 0; i < 100; i += 1 {
		go log.Println(i) // this is thread safe
	}

	for i := 0; i < 100; i += 1 {
		go fmt.Println(i) // this is not thread safe
	}

	log.SetOutput(ioutil.Discard)

	log.Println("Entry 2")

	defer log.Println("Will not be logged")

	log.Fatal("Exit") // fatal calls an system exist so any defers would not have time to be called in a program that logs fatal
}
