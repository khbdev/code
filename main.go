package main

import (
	"fmt"
	"sync"


)


type Job struct {
	ID int
	Name string
}



func worker(id int, jobs <-chan Job, wg *sync.WaitGroup){
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d is processing job %d and name  %s\n", id, job.ID, job.Name)
	}
}


func main(){
	numberWorker := 4
	numberJobs := 10

	jobs := make(chan Job, numberJobs)

	var wg sync.WaitGroup

	wg.Add(numberWorker)
	for i := 0; i < numberWorker; i++ {
		go worker(i, jobs, &wg)
	}

	for i := 0; i < numberJobs; i++ {
		jobs <- Job{ID: i, Name: "Azizbek"}
	}
	close(jobs)

	wg.Wait()
}