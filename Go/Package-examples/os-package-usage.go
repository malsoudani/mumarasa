package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	wd, err := os.Getwd() // get current working directory
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(wd)

	fmt.Printf(os.Getenv("PATH")) // get an an environment variable

	currUser, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(currUser.Username)

	cmd := exec.Command("ls", "-ltr")

	stdErr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(stdErr))
}
