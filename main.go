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
		ch1 <- "Salom, Azizbek! ch1dan"
	}()
		go func() {
		time.Sleep(1 *time.Second)
		ch2 <- "Qonday, Azizbek! ch2dan"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Channel 1dan ", msg1)
	case msg2 := <-ch2:
		fmt.Println("Channel 2dan ", msg2)
	case <-time.After(3 *time.Second):
		fmt.Println("Timeout: hech narsa kelmadi")
	}	

	fmt.Println("ish tugadi")
}