package main

import (
	"fmt"
	"time"
)


func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)


	go func() {
		time.Sleep(1 *time.Second)
		ch1 <- "hello one channel"
	}()

	go func() {
		time.Sleep(2 *time.Second)
		ch2 <- "Hello two channel"
	}()

	for i := 0; i < 2; i++ {
		select {
		case val1 := <- ch1:
			fmt.Println(val1)
		case val2 := <- ch2:
			fmt.Println(val2)
		}
	}
}