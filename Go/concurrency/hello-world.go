package main

import (
	"fmt"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
}

func main() {
	hello, world := new(Job), new(Job)

	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	hello.text = "world"
	hello.i = 0
	hello.max = 5

	go outputText(hello)
	outputText(world)
}
