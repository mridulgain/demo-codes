package main

import (
	"fmt"
	"sync"
	"testing"
)

func writer(wg *sync.WaitGroup) {
	fmt.Println("Start Producing")
	for i := 0; i < 10; i++ {
		// buffer <- i
		fmt.Println(i)
	}
	fmt.Println("End Producing")
	wg.Done()
}

func Test_wg(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go writer(&wg)
	go writer(&wg)
	wg.Wait()
}
