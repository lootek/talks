package main

import (
	"fmt"
	"io"
)

// START OMIT
type reader struct {
	cnt int
}

func (r *reader) Read(b []byte) (int, error) {
	r.cnt++ // panic // HL
	return 0, nil
}

func main() {
	var rr *reader
	var r3 io.Reader = rr

	n, err := r3.Read([]byte("foo"))
	fmt.Println(n, err)
}

// END OMIT
