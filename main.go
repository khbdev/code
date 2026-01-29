package main

import (
	"fmt"
	"sync"
)


func main(){

var wg sync.WaitGroup


	ch := make(chan  int)
 wg.Add(2)
	go func() {
		ch <- 1
		wg.Done()
	}()

	go func() {
		num := <- ch
		fmt.Println(num)
		wg.Done()
	}()
	wg.Wait()
}