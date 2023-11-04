package main

import (
	"fmt"
	"sync"
)

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
	shutdown    chan struct{}
	waitGroup   sync.WaitGroup
}

func NewWorkerPool(noWorkers, chanBuffer int) *workerPool {
	wp := &workerPool{
		maxWorker:   noWorkers,
		queuedTaskC: make(chan func(), chanBuffer),
		shutdown:    make(chan struct{}),
	}
	return wp
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		wp.waitGroup.Add(1)
		go func() {
			defer wp.waitGroup.Done()
			for {
				select {
				case task, ok := <-wp.queuedTaskC:
					if !ok {
						return
					}
					task()
				case <-wp.shutdown:
					return
				}
			}
		}()
	}
}

func (wp *workerPool) AddTask(task func()) {
	select {
	case wp.queuedTaskC <- task:
	default:
		// Handle task queuing failure if the channel is full
		fmt.Println("Task dropped: queue full")
	}
}

func (wp *workerPool) Shutdown() {
	close(wp.queuedTaskC)
	close(wp.shutdown)
	wp.waitGroup.Wait()
}

func main() {
	wp := NewWorkerPool(10, 10)
	wp.Run()

	for i := 0; i < 1000; i++ {
		wp.AddTask(func() {
			// Do something
		})
	}

	// Graceful shutdown
	wp.Shutdown()
}
