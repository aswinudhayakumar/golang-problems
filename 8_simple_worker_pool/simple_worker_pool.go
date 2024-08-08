package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	newWorkerPool(5, 10, taskFunction)
	fmt.Println("total time taken - ", time.Since(t))
}

func taskFunction() {
	fmt.Println("task started")
	time.Sleep(1 * time.Second)
}

func runner(jobs chan func(), doneCh chan struct{}) {
	for task := range jobs {
		task()
		doneCh <- struct{}{}
	}
}

func newWorkerPool(maxNoOfWorkers, noOfJobs int, task func()) {
	jobs := make(chan func(), noOfJobs)
	doneCh := make(chan struct{})

	for i := 0; i < maxNoOfWorkers; i++ {
		go runner(jobs, doneCh)
	}

	for j := 0; j < noOfJobs; j++ {
		jobs <- task
	}
	close(jobs)

	for k := 0; k < noOfJobs; k++ {
		<-doneCh
		fmt.Println("done")
	}
}
