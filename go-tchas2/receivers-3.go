package main

import "fmt"

type number struct {
	i int
}

// START OMIT
func (n *number) inc() {
	n.i++ // HL
}

func (n *number) print() {
	fmt.Println(n) // HL
}

func main() {
	n1 := &number{} // HL
	fmt.Println("calling n1.print()...")
	n1.print()
	fmt.Println("calling n1.inc()...")
	n1.inc()

	var n2 *number // HL
	fmt.Println("calling n2.print()...")
	n2.print()
	fmt.Println("calling n2.inc()...")
	n2.inc() // panic // HL
}

// END OMIT
