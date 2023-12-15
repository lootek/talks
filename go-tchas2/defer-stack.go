package main

import (
	"fmt"
)

// START OMIT
func f1() {
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func main() {
	defer f1()
	defer f2()

	fmt.Println("main")
}

// END OMIT
