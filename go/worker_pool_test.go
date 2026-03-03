package main

import (
	"log"
	"testing"
)

func Test_worrkerPool(t *testing.T) {
	// creare some dummy tasks
	dummyTasks := []Task{}
	for i := 0; i < 20; i++ {
		dummyTasks = append(dummyTasks, Task{id: i + 1})
	}
	// init a worker pool with
	workerPool := WorkerPool{
		concurrency: 3,
		tasks:       dummyTasks,
	}
	workerPool.Run()
	log.Printf("All worker has finished\n")
}
