package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func worker(ctx context.Context, id int) {
	done := make(chan bool)
	fmt.Println("Worker", id, ": Bắt đầu công việc.")

	go func() {
		fmt.Printf("Worker %d: Đang thực hiện công việc...\n", id)
		wait := time.Duration(rand.Intn(10)) * time.Second
		fmt.Printf("Thoi gian worker %d thuc hien la %s\n", id, wait)
		time.Sleep(wait)

		done <- true
	}()

	select {
	case <-done:
		fmt.Printf("Worker %d: Công việc hoàn thành.\n", id)
	case <-ctx.Done():
		fmt.Printf("Worker %d: Context đã hết thời gian, hủy bỏ công việc.\n", id)
	}
}

func main() {
	now := time.Now()
	deadline := time.Now().Add(5 * time.Second)
	fmt.Println(now, deadline, "deadline context\n")
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i)
	}

	time.Sleep(10 * time.Second)
}
