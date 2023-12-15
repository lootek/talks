package main

import (
	"fmt"
)

// START OMIT
func main() {
	s := []int{}

	for i := 0; i < 20; i++ {
		fmt.Printf("%p %d %d %v\n", s, len(s), cap(s), s)
		s = append(s, i)
	}

	for i := 20; i < 2500; i++ {
		fmt.Printf("%p %d %d\n", s, len(s), cap(s))
		s = append(s, i)
	}

	// starting cap    growth factor
	// 256             2.0
	// 512             1.63
	// 1024            1.44
	// 2048            1.35
	// 4096            1.30
}

// END OMIT
