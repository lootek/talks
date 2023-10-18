package main

import (
	"os"
	"regexp"
)

// START OMIT
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	return digitRegexp.Find(b) // HL
}

func CopyDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = digitRegexp.Find(b)   // HL
	c := make([]byte, len(b)) // HL
	copy(c, b)                // HL
	return c
}

// END OMIT
