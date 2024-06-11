// use buffer channel to process a line of working
// write func to calculate the square of a number
package main

import (
	"fmt"
	"time"
)

const numOfworkers int64 = 5

func main() {
	// array init
	numArray := []int64{1, 3, 5, 6, 8, 20, 40, 60, 80, 100}

	queue := initQueue(numArray)
	var i int64
	for i = 1; i <= numOfworkers; i++ {
		go square(queue, i)
	}
	time.Sleep(time.Minute * 5)

}

func square(queue <-chan int64, name int64) {
	for v := range queue {
		fmt.Printf("Worker %d is processing number %d . Resutlt %d \n", name, v, v*v)
		time.Sleep(time.Second)
	}
}

func initQueue(numberSlice []int64) <-chan int64 {
	numberOfJobs := len(numberSlice)
	queue := make(chan int64, 100)
	go func() {
		for i := 0; i < numberOfJobs; i++ {
			queue <- numberSlice[i]
		}
		close(queue)
	}()
	return queue
}
