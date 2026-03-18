package main

import (
	"fmt"
	"sync"
)


func main(){
	var wg sync.WaitGroup


for i := 0; i < 1000; i++ {
	wg.Add(1)

	go func(id int) {
		defer wg.Done()
		fmt.Printf("gorutina id  %d\n", id)
	}(i)

	wg.Wait()


}
	fmt.Println("Gorutinlar ishini tugatdi")
}