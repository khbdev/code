package main



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

