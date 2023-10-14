package main

import (
	"errors"
	"fmt"
)

func f1() error {
	return nil
}

func f2() error {
	return nil
}

func f3() error {
	return nil
}

// START OMIT
func F() error {
	err := f1()
	if err != nil {
		return err
	}

	err = f2()
	if err != nil {
		return err
	}

	var errs error
	for i := 0; i < 10; i++ {
		err := f3() // HL
		if err != nil {
			errs = errors.Join(errs, err) // HL
			continue
		}
	}

	return errs
}

// END OMIT

func main() {
	fmt.Println(F())
}
