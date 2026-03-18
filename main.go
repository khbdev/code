package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)


func main(){
	var wg sync.WaitGroup

fmt.Println("Boshlanishda Gorutinlar soni: ", runtime.NumGoroutine())
for i := 0; i < 1000; i++ {
	wg.Add(1)

	go func(id int) {
		defer wg.Done()
		fmt.Printf("gorutina id  %d\n", id)
	}(i)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Ishlayotganda goroutine soni:", runtime.NumGoroutine())

	wg.Wait()


}
	fmt.Println("Gorutinlar ishini tugatdi")
		fmt.Println("Tugagandan keyin goroutine soni:", runtime.NumGoroutine())
}