package main

import (
	"fmt"
	"sync"
)



func worker(id int, jobs <- chan int, result chan <- int, wg *sync.WaitGroup){
	defer wg.Done()
	for  v := range jobs {
		fmt.Printf("Worker Id %d, job %d\n", id, v)
		result <- 2 *  v
	}
}

func main(){
	jobs := make(chan int, 10)
	result := make(chan int, 10)


	jobsList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			worker(id, jobs, result, &wg)
		}(i)
	}

	for _, v := range jobsList {
	 jobs <- v
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(result)
	}()

	for  res := range result {
		fmt.Println("Natija: ", res)
	}

	
}