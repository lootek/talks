package main

// START OMIT
func main() {
	const N = 20

	s1 := make([]int, N) // HL
	for i := 0; i < N; i++ {
		s1[i] = i
	}

	s2 := make([]int, 0, N) // HL
	for i := 0; i < N; i++ {
		s2 = append(s2, i)
	}
}

// END OMIT
