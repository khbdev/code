package main

import "fmt"



type User struct {
	ID int 
	Name string
}


var (
	users []User
	lastID = 0
)

func CreateUser(name string) User {
	lastID++
	user := User {
      ID: lastID,
	  Name: name,
	}
	return  user
}

func main(){
	var name string

	for i := 1; i <= 5; i++ {
		fmt.Print("Name: ", )
	fmt.Scanln(&name)
	}

	
}