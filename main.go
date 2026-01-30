package main

import (
	"fmt"
	"sync"

)


func worker(id int, jobs <- chan int, result chan<- int, wg *sync.WaitGroup) {
	for  job := range jobs {
			fmt.Printf("Worker %d job %d ni ishlayapdi\n", id, job)
		
		result <- job * 2 
		 wg.Done()
	}
}

func main(){
	jobs := make(chan int, 10)
	result := make(chan int, 10)
	value := 10

var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		go func(id int) {
			worker(id, jobs, result, &wg)
		}(i)
	}

	

	for i := 1; i <= value; i++ {
		wg.Add(1)
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= value; i++ {
	fmt.Println("Malumot", <-result)
	}
	wg.Wait()
	close(result)
}