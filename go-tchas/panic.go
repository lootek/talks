package main

import (
	"fmt"
	"time"
)

// START OMIT
func f1() {
	time.Sleep(time.Second)
	panic("boo") // HL
}

func main() {
	defer func() {
		if r := recover(); r != nil { // HL
			fmt.Println("recovered from:", r)
		}
	}()

	go f1() // HL

	// do actual stuff
	time.Sleep(time.Minute)
}

// END OMIT
