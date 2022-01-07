package workerpool

import (
	"context"
	"myhttp/pkg/httpRequester"
	"sync"
)

//NewWorkerPool will initialize the worker pool with max number of workers that we want
func (wp *workerPool)NewWorkerPool(maxWorker int){
		wp.maxWorker =  maxWorker
		wp.wg =     sync.WaitGroup{}
}


//AddTask will append task tp queuedTask fields
func (wp *workerPool) AddTask(tasks []string) {
	wp.queuedTask = make(chan string,len(tasks))
	for _, task := range tasks{
		wp.queuedTask <- task
	}
	close(wp.queuedTask)
}

//Run will start the workers and call the MakeRequest method
func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for url := range wp.queuedTask{
				if err := httpRequester.HTTPRequester.MakeRequest(context.Background(),url);err != nil {
					break
				}
			}
		}()
	}
	wp.wg.Wait()
}

