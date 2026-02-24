package main

import (
	"fmt"
	"os"
	"time"
)



func main(){
	fmt.Println("My pid: ", os.Getpid())
	fmt.Println("Parent PID: ", os.Getppid())

	for i := 0; i < 50; i++ {
		fmt.Println("Alice second: ", i)
		time.Sleep(1 * time.Second)
	}
}