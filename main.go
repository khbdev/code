package main

import (
	"fmt"
	"time"
)


func worker(id int, jobs <-chan int){
	fmt.Printf("Worker id %d, job %d ni ishlayabdi", id, <-jobs)
	time.Sleep(1  * time.Second)
}


func main(){
	jobs := make(chan int)

	for i := 1; i < 3; i++ {
		go func() {
			worker(i, jobs)
		}()
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	time.Sleep(5 *time.Second)
}


