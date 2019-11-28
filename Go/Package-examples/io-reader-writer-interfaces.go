package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type lowerCaseReader struct {
	text string
}

func NewLowerCaseReader(text string) *lowerCaseReader {
	return &lowerCaseReader{text: text}
}

func (r *lowerCaseReader) Read(p []byte) (int, error) {
	buf := make([]byte, len(r.text))

	for i := 1; i < len(buf); i++ {
		buf[i] = r.text[i] | 0x20
	}

	n := copy(p, buf)
	return n, io.EOF

}

func main() {

	r := NewLowerCaseReader("ALL CAPITAL LETTERS BLAH")
	resp, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("something was wrong with converting to lowercase")
	}

	fmt.Println(string(resp))
}
