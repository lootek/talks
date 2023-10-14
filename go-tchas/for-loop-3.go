package main

import (
	"fmt"
	"time"
)

func main() {

	// START OMIT
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	// END OMIT

	time.Sleep(time.Second)
}
