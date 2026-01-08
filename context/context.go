package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"time"
)

func main() {
	ctx, cF := context.WithTimeout(context.Background(), 3*time.Second)
	defer cF()

	go func() {
		for {
			fmt.Println("time")
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-ctx.Done(): // Context timed out or was canceled
		fmt.Println("Operation timed out or was canceled.")
		return
	}

	time.Sleep(100 * time.Second)
}
