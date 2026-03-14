package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

// multiple producer multiple consumer
func producerr(id int, n int, odd chan int) {
	for i := 0; i < n; i++ {
		// Just to make the data unique to the producer
		val := (id * 1000) + i
		odd <- val
	}
	log.Printf("Producer %d finished sending %d items", id, n)
}

func consumerr(id int, wg *sync.WaitGroup, odd chan int) {
	defer wg.Done()
	log.Println("Consumer id starting: ", id)
	for i := range odd {
		// Simulate some work
		_ = fmt.Sprintf("processed: %d by consumer %d", i, id)
		log.Printf("Consumer %d processed: %d", id, i)
	}
}

// multiple producer & consumers
func Test_PC_Multi(t *testing.T) {
	producerWG := sync.WaitGroup{}
	consumerWG := sync.WaitGroup{}

	myChan := make(chan int, 10)
	itemsPerProducer := 50
	numProducers := 2
	numConsumers := 3

	// 1. Start Multiple Producers
	for i := 1; i <= numProducers; i++ {
		producerWG.Add(1)
		go func(pID int) {
			defer producerWG.Done()
			producerr(pID, itemsPerProducer, myChan)
		}(i)
	}

	// 2. Start Multiple Consumers
	for i := 1; i <= numConsumers; i++ {
		consumerWG.Add(1)
		go consumerr(i, &consumerWG, myChan)
	}

	// 3. The "Closer" Goroutine
	// Crucial: It monitors all producers and shuts down the pipe
	go func() {
		producerWG.Wait()
		close(myChan)
		log.Println("All producers finished. Channel closed.")
	}()

	// 4. Wait for all consumers to drain the channel
	consumerWG.Wait()
	log.Println("All work complete.")
}

// go test -v -run Test_PC_Multi
