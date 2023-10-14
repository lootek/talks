package main

import (
	"os"
	"regexp"
)

// START OMIT
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	return digitRegexp.Find(b)
}

func CopyDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// END OMIT
