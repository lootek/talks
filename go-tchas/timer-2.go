package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.TODO())
	defer cancelFn()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker")
			return
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		}
	}
}
