package main

import (
	"fmt"
)

// START OMIT
func f() (err error) {
	err = fmt.Errorf("new error")
	defer func() {
		err = fmt.Errorf("more context: %w", err)
	}()
	return err
}

func main() {
	fmt.Println(f())
}

// END OMIT
