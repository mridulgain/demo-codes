package main

import (
	"fmt"
	"testing"

	"sync"
)

func worker(nums <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range nums {
		if n%2 == 1 {
			results <- n
		}
	}
}

func oddNumbers(n int, workers int) []int {
	results := make(chan int)
	nums := make(chan int)
	var wg sync.WaitGroup
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(nums, results, &wg)
	}
	//producer
	go func() {
		for i := 0; i <= n; i++ {
			nums <- i
		}
		close(nums)
	}()

	// consumer
	go func() {
		wg.Wait()
		close(results)
	}()

	var res []int

	for r := range results {
		res = append(res, r)
	}
	return res

}

func Test_ProducerConsumer(t *testing.T) {
	res := oddNumbers(20, 3)
	fmt.Println(res)
}
