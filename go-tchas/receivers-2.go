package main

import "fmt"

// START OMIT
type number struct {
	i int
}

func (n *number) inc() { // HL
	n.i++
}

func main() {
	n := number{8}

	fmt.Printf("%#v\n", n)
	n.inc()
	fmt.Printf("%#v\n", n)
}

// END OMIT
