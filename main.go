package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(rb []byte) (int, error) {
	for i := 0; i < len(rb); i++ {
		rb[i] = 'A'
	}
	return len(rb), nil
}

func main() {
	reader.Validate(MyReader{})
}
