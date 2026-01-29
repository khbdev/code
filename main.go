package main

import (
	"fmt"
	"time"
)


func main(){

for i := 0; i < 4; i++ {
	go func() {
		fmt.Println("Hello")
	}()
	time.Sleep(1 * time.Second)
}
time.Sleep(2 * time.Second)

}