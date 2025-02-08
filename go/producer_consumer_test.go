package main

import (
	"fmt"
	"sync"
	"testing"
)

func producer(buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("produced: ", i*2)
		buffer <- i * 2
	}
	close(buffer)
}

func consumer(buffer chan int, wg *sync.WaitGroup) {
	wg.Wait()
	for i := range buffer {
		fmt.Println("consumed: ", i)
	}
}

func Test_pc(t *testing.T) {
	buf := make(chan int, 10)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go producer(buf, wg)
	// read from channel
	go consumer(buf, wg)
	wg.Wait()
}
