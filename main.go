package main

import "fmt"



func main(){
	ch := make(chan int)


	for i := 1; i <= 3; i++ {
		go func() {
			ch <- i
		}()
	}
	close(ch)

	sam := <- ch
	fmt.Println(sam)

	
}