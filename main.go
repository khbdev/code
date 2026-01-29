package main

import (
	"fmt"
	"time"
)


func main(){

	gortina := 5

	for i := 0; i < gortina; i++ {
		go func() {
			fmt.Println("Salom")
		}()
		time.Sleep(1 * time.Second)
	}
}