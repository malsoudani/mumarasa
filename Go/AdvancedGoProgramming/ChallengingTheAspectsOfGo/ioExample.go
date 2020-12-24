package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "blah")
	fmt.Fprint(w, "some other blah")

	// always flush your buffered writer once you are done!
	w.Flush()

	intScanner := bufio.NewScanner(strings.NewReader("1,2,3,4,5,6,7,8"))

	intScanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF {
			if len(data) == 0 {
				return 0, nil, nil
			}
			return len(data), data, nil
		}

		if i := strings.IndexRune(string(data), ','); i >= 0 {
			return i + 1, data[0:i], nil
		}
		return 0, nil, nil
	})

	buf := make([]byte, 2)

	// scan 2 bytes at a time
	intScanner.Buffer(buf, bufio.MaxScanTokenSize)
	for intScanner.Scan() {
		fmt.Printf("%s", intScanner.Text())
	}
}
