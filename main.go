package main

import (
	// "fmt"
	"time"
)


func Worker(name string){
	// for i := 1; i <= 3; i++ {
	//   fmt.Println(name, "Workering", i)
	//   time.Sleep(time.Second)
	// }
	panic("boom")
}


func main(){
	go Worker("Thread-1")
	go Worker("Thread-2")

	time.Sleep(4 * time.Second)
}