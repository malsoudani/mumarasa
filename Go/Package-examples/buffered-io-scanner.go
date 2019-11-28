package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(w, "Blah Blah")
	fmt.Fprintf(w, "asdasdasdasda \n")
	w.Flush()

	numberInput := "1,2,3,4,5,6,4,3,2,2,3,4,5,6,6,3,2,2,3,4,5,5,6,5,4,3,2,9,6,7,7,6,6,6,6,5"

	intScanner := bufio.NewScanner(strings.NewReader(numberInput))

	// custom splitter based on commas

	splitByComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) { // this signature is specific to the dispatcher of the method which is intScanner.Split(...) in this example
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := strings.IndexRune(string(data), ','); i >= 0 {
			return i + 1, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return
	}

	intScanner.Split(splitByComma) // splitByComma is supposed to be a closure predefined

	buf := make([]byte, 2)

	intScanner.Buffer(buf, bufio.MaxScanTokenSize)

	for intScanner.Scan() {
		fmt.Printf(intScanner.Text())
	}

}
