package main

import (
	"fmt"
	"sync"
)



type User struct {
	ID int 
	Name string
}


var (
	users []User
	lastID = 0
	wg sync.WaitGroup
	mu sync.Mutex
)

func CreateUser(name string, wg *sync.WaitGroup) User {
	wg.Done()
	mu.Lock()
	lastID++
	user := User {
      ID: lastID,
	  Name: name,
	}
	users = append(users, user)
	
	mu.Unlock()
	return  user
}

func main(){
	var name string
	
  
	for i := 1; i <= 5; i++ {
		fmt.Print("Name: ", )
	fmt.Scanln(&name)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Print("Name: ", )
	fmt.Scanln(&name)
		}(i)
	}
	wg.Wait()
}