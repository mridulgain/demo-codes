package main

import (
	"log"
	"sync"
	"time"
)

// define task
type Task struct {
	id int
}

// define how to process that task
func (t *Task) process() {
	log.Printf("\tProcessing task with id: %d\n", t.id)
	// simulating task process
	time.Sleep(2 * time.Second)
}

// define worker pool
type WorkerPool struct {
	tasks       []Task
	concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

// methods on workerpool
func (wp *WorkerPool) worker(id int) {
	for t := range wp.taskChan {
		log.Printf("task %d processed by worker %d\n", t.id, id)
		t.process()
	}
	wp.wg.Done()
}

func (wp *WorkerPool) Run() {
	wp.taskChan = make(chan Task, len(wp.tasks))
	wp.wg = sync.WaitGroup{}
	wp.wg.Add(wp.concurrency)
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker(i + 1)
	}
	for _, j := range wp.tasks {
		wp.taskChan <- j
	}
	close(wp.taskChan)
	wp.wg.Wait()
}
