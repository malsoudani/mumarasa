package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
		}

		cmdString = strings.TrimSuffix(cmdString, "\n")
		cmdStringsArr := strings.Fields(cmdString)
		cmd := exec.Command(cmdStringsArr[0], cmdStringsArr[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println(os.Stderr, err)
		}
	}
}
