package main

import (
	"fmt"
	"time"
)

func main() {

	// START OMIT
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// END OMIT

	time.Sleep(time.Second)
}
