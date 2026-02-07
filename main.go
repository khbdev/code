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

func CreateUser(name string, wg *sync.WaitGroup){
defer	wg.Done()
	mu.Lock()
	defer 	mu.Unlock()
	lastID++
	user := User {
      ID: lastID,
	  Name: name,
	}
	users = append(users, user)
}

func getUsers(id int) User {
	user := users[id]
	return  user
}

func main(){
	names := make([]string, 0, 5)
	
  
	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Name: ", )
	fmt.Scanln(&name)
	names = append(names, name)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			CreateUser(names[i], &wg)
		
		}(i)
	}
		
	wg.Wait()
	fmt.Println(users)
  
	getUsers(1)



}