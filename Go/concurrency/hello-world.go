package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, goGroup *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	goGroup.Done() // denoting a single iteration
}

func main() {
	goGroup := new(sync.WaitGroup)
	hello, world := new(Job), new(Job)

	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	world.text = "world"
	world.i = 0
	world.max = 5

	go outputText(hello, goGroup)
	go outputText(world, goGroup)

	goGroup.Add(2) // the Add method here specifies how many Done() messages should the goGroup recieve before satisfying its wait
	goGroup.Wait() // this does the actual waiting
}
