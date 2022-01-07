package main

import (
	"fmt"
	"log"
	"myhttp/config"
	"myhttp/pkg/httpRequester"
	"myhttp/pkg/workerpool"
	"net/http"
	"os"
	"strconv"
	"time"
)

func init(){
	dir, _ := os.Getwd()
	if err := config.Confs.Load(dir + "/config.yaml"); err != nil {
		log.Println(err)
		return
	}
	//initialize httpRequester
	httpRequester.HTTPRequester.NewHttpRequester(http.Client{})
}


func main() {
	// start of program
	start := time.Now()
	// find out the number of goroutines and Urls
	numGoroutines := 10
	Urls := os.Args[1:]

	if os.Args[1] == "-parallel"{
		numGoroutines ,_= strconv.Atoi(os.Args[2])
		Urls = os.Args[3:]
	}

	// Initialize the worker pool and start to run ...
	workerpool.Worker.NewWorkerPool(numGoroutines)

	workerpool.Worker.AddTask(Urls)

	workerpool.Worker.Run()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}


