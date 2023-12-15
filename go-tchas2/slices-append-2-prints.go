package main

import (
	"fmt"
)

// START OMIT
func main() {
	const N = 20

	s1 := make([]int, N) // HL
	fmt.Printf("(%p) len %d, cap %d, %v\n", s1, len(s1), cap(s1), s1)
	for i := 0; i < N; i++ { // HL
		fmt.Printf("(%p) len %d, cap %d, %v\n", s1, len(s1), cap(s1), s1)
		s1[i] = i // HL
	} // HL
	fmt.Printf("(%p) len %d, cap %d, %v\n", s1, len(s1), cap(s1), s1)
	fmt.Println()

	s2 := make([]int, 0, N) // HL
	fmt.Printf("(%p) len %d, cap %d, %v\n", s2, len(s2), cap(s2), s2)
	for i := 0; i < N; i++ { // HL
		fmt.Printf("(%p) len %d, cap %d, %v\n", s2, len(s2), cap(s2), s2)
		s2 = append(s2, i) // HL
	} // HL
	fmt.Printf("(%p) len %d, cap %d, %v\n", s2, len(s2), cap(s2), s2)
}

// END OMIT
