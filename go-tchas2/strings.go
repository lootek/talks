package main

import "fmt"

func main() {

	// START OMIT
	s := "hêllo"
	fmt.Printf("len=%d\n", len(s))

	for i := range s { // HL
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Println()

	for i, r := range s { // HL
		fmt.Printf("position %d: %c\n", i, r)
	}
	fmt.Println()

	for i := 0; i < len(s); i++ { // HL
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Println()

	runes := []rune(s)
	for i, r := range runes { // HL
		fmt.Printf("position %d: %c\n", i, r)
	}

	// END OMIT
}
