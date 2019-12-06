package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	args := "-ltr"
	cmd := exec.Command("ls", strings.Split(args, " ")...)

	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}
