package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10
	jobs := make(chan int) // The "Buffer"
	var wg sync.WaitGroup

	// 1. START CONSUMERS (Workers)
	// These wait for data to arrive on the 'jobs' channel
	numberOfWorkers := 3
	for w := 1; w <= numberOfWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for num := range jobs {
				fmt.Printf("Worker %d processed odd number: %d\n", workerID, num)
			}
		}(w)
	}

	// 2. START PRODUCER
	// This finds the numbers and sends them to the channel
	go func() {
		for i := 1; i <= n; i++ {
			if i%2 != 0 {
				jobs <- i
			}
		}
		close(jobs) // Crucial: tell consumers no more data is coming
	}()

	wg.Wait()
}
