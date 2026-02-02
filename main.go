package main

import (
	"fmt"
	"sync"
)


var (
	number int
	mu sync.Mutex
	wg sync.WaitGroup
)


func incrament(){
	defer wg.Done()
	mu.Lock()
	number++
	mu.Unlock()
}


func main(){
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrament()
	}
	wg.Wait()
	fmt.Println(number)
}