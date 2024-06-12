// use buffer channel to process a line of working
// write func to calculate the square of a number
package main

import (
	"fmt"
	"sync"
)

const numOfworkers int64 = 5

func main() {
	// array init
	numArray := []int64{1, 3, 5, 6, 8, 20, 40, 60, 80, 100}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	queue := initQueue(wg, numArray)
	var i int64
	for i = 1; i <= numOfworkers; i++ {
		wg.Add(1)
		go square(wg, queue, i)
	}
	wg.Wait()
}

func square(wg *sync.WaitGroup, queue <-chan int64, name int64) {
	defer wg.Done()
	for v := range queue {
		fmt.Printf("Worker %d is processing number %d . Resutlt %d \n", name, v, v*v)
	}
}

func initQueue(wg *sync.WaitGroup, numberSlice []int64) <-chan int64 {
	defer wg.Done()
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
