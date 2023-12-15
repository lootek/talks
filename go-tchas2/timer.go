package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.TODO())
	defer cancelFn()

	// START OMIT
	go func() {
		time.Sleep(time.Second)
		cancelFn()
	}()

	for {
		select {
		case <-time.After(time.Minute):
			fmt.Println("timed out")
			break
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		}
	}
	// END OMIT
}
