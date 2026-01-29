package main

import (
	"fmt"
	"sync"
)

func GoruOne(){
	fmt.Println("Salom")
}


func main(){

var wg sync.WaitGroup

wg.Add(1)
go func() {
	wg.Done()
	GoruOne()
}()

wg.Wait()



}