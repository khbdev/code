package main

import (
	"fmt"
	"sync"
)


func main(){
	total := 1_000_000
	worker := 10

	chunk := total / worker

	results := make([]int, worker)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			start := id * chunk
			end := start + chunk

			local := 0
			for i := start; i < end; i++ {
				local++
			}
			results[id] = local
		}(i)
	}
	wg.Wait()

	sum := 0
	for _, v := range results {
		sum+=v
	}
	fmt.Println("Natija: ", sum)
}