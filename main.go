package main

type User struct {
	Name string
	Age  int
}

var (
	users  = []User
	lastId = 0
)

func CreateUser(name string, age int) {
	lastId++
	user := User{
		Name: name,
		Age:  age,
	}
	users = append(users, user)

}

func GetUsers() User {
	return users
}
func getByIdUsers(id int) User {
	user := users[id]

	return user
}

func main() {
	CreateUser("Gopher", 17)
	GetUsers()
	getByIdUsers(1)

}
