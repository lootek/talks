package main

import (
	"fmt"
)

// START OMIT
func main() {
	s := []int{}

	for i := 0; i < 20; i++ {
		fmt.Println(len(s), cap(s), s)
		s = append(s, i)
	}
}

// END OMIT
