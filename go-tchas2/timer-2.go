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

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("timeout")
			return
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		}
	}
	// END OMIT
}
