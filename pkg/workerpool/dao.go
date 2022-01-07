package workerpool

import (
	"sync"
)

var (
	Worker workerPoolInterface = &workerPool{}
)

type workerPoolInterface interface {
	NewWorkerPool(maxWorker int)
	Run()
	AddTask(tasks []string)
}

type workerPool struct {
	maxWorker   int
	queuedTask  chan string
	wg sync.WaitGroup
}



