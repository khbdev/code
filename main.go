package main

import "fmt"

type User struct {
	Name string
	Age int
}


var users []User

func main(){
for i := 0; i < 1_000_000_000; i++ {
		users = append(users, User{Name: "Azizbek", Age: 17})
}

for _, v := range users {
	fmt.Println(v.Name)
	fmt.Println(v.Age)
}
}