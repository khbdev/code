package main

import "fmt"

type User struct {
	Id   int
	Name string
}

var users []User
var lastId int = 0

func AddliSt(name string) {
	lastId++
	user := User{
		Id:   lastId,
		Name: name,
	}
	users = append(users, user)
	fmt.Println("Create SuccessFull")
}
func ListUsers() {
	fmt.Println(users)
}
func getById(id int) (User, bool) {
	for _, user := range users {
		if user.Id == id {
			return user, true
		}
	}
	return User{}, false
}
func main() {
	AddliSt("Azizbek")
	ListUsers()
	u, ok := getById(1)
	if !ok {
		fmt.Println("Topilmadi")
		return
	}
	fmt.Println(u)
}
