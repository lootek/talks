package main

import (
	"strconv"
	"strings"
	"testing"
)

// START OMIT
func BenchmarkStringConcat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += strconv.Itoa(i)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
}

// BenchmarkStringConcat-8      401053   147346 ns/op   1108686 B/op   2 allocs/op
// BenchmarkStringBuilder-8   29307392       45 ns/op        52 B/op   0 allocs/op

// END OMIT
