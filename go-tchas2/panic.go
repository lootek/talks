package main

import (
	"fmt"
	"time"
)

// START OMIT
func f1() {
	defer func() {
		if r := recover(); r != nil { // HL
			fmt.Println("recovered from:", r)
		}
	}()

	time.Sleep(time.Second)
	panic("boo")
}

func main() {
	go f1() // HL

	// do actual stuff
	time.Sleep(time.Minute)
}

// END OMIT
