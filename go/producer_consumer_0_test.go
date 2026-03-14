package main

import (
	"log"
	"sync"
	"testing"
)

// one producer multiple consumer
func producer(n int, odd chan int) {
	for i := 0; i < n; i++ {
		odd <- i*2 + 1
	}
}

func consumer(id int, wg *sync.WaitGroup, odd chan int) {
	log.Println("Condumer id: ", id)
	for i := range odd {
		log.Println("processed: ", i)
	}
	wg.Done()
}

func Test_PC(t *testing.T) {
	producerWG := sync.WaitGroup{}
	consumerWG := sync.WaitGroup{}

	myChan := make(chan int, 10) // Small buffer for backpressure
	count := 100
	concurrency := 3

	// 1. Start Producers
	producerWG.Add(1)
	go func() {
		defer producerWG.Done()
		producer(count, myChan)
	}()

	// 2. Start Consumers
	for i := 0; i < concurrency; i++ {
		consumerWG.Add(1)
		go consumer(i+1, &consumerWG, myChan)
	}

	// 3. The "Closer" Goroutine
	// This waits for producers in the background so it doesn't block consumers
	go func() {
		producerWG.Wait()
		close(myChan)
	}()

	// 4. Wait for work to be done
	consumerWG.Wait()
}

// go test -v -run Test_PC
