package main

import (
	"fmt"
	"time"
)



func GetUser() string{
	time.Sleep(1 *time.Second)
	return  "User: Azizbek"
}

func GetProduct() string {
	time.Sleep(1 *time.Second)
	return  "Product: Macbook"
}


func main(){
	start := time.Now()


	userChan := make(chan string)
	procutChan := make(chan string)

	go func() {
		userChan <- GetUser()
	}()
	go func() {
		procutChan <- GetProduct()
	}()

	user := <-userChan
	product := <-procutChan

	fmt.Println(user)
	fmt.Println(product)
	fmt.Println("Otkan vaqt: ", time.Since(start))
}