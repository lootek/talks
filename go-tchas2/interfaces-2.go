package main

import (
	"fmt"
	"io"
)

// START OMIT
type reader struct{}

func (r *reader) Read(_ []byte) (int, error) {
	// TODO: implement me // HL
	return 0, nil // HL
}

func main() {
	var rr *reader
	var r3 io.Reader = rr

	n, err := r3.Read([]byte("foo"))
	fmt.Println(n, err)
}

// END OMIT
