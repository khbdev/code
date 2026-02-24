package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)



func Worker(id int, wg *sync.WaitGroup){
	defer wg.Done()

for i := 0; i < 3; i++ {
	fmt.Println("Worker", id, "iteration", i)
		time.Sleep(200 * time.Millisecond)
}
}


func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Cpu Cores:", runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}

	wg.Wait()
}