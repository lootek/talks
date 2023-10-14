package main

import (
	"os"
)

// START OMIT
func main() {
	for _, filename := range []string{"file1.txt", "file2.txt", "file3.txt"} {
		f, _ := os.Open(filename)
		defer f.Close()

		// read the file ...
		// do some stuff ...
	}
}

// END OMIT
