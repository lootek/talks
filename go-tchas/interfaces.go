package main

import (
	"fmt"
	"io"
)

// START OMIT
type reader struct{}

func (r *reader) Read(_ []byte) (int, error) {
	// TODO: implement me
	return 0, nil
}

func main() {
	var r1 io.Reader
	fmt.Printf("r1: %#v (%v)\n", r1, r1 == nil)

	r2 := io.Reader(nil)
	fmt.Printf("r2: %#v (%v)\n", r2, r2 == nil)

	var rr *reader
	var r3 io.Reader = rr
	fmt.Printf("r3: %#v (%v)\n", r3, r3 == nil)
}

// END OMIT
