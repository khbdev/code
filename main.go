package main

import (
	"errors"
	"fmt"
)

// 🔹 User struct
type User struct {
	ID    int
	Name  string
	Email string
}

// 🔹 In-memory storage
var users = make(map[int]User)
var lastID = 0

// 🔹 Create
func CreateUser(name, email string) User {
	lastID++
	user := User{
		ID:    lastID,
		Name:  name,
		Email: email,
	}
	users[user.ID] = user
	return user
}

// 🔹 Get
func GetUser(id int) (User, error) {
	user, exists := users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// 🔹 Update
func UpdateUser(id int, name, email string) (User, error) {
	user, err := GetUser(id)
	if err != nil {
		return User{}, err
	}
	user.Name = name
	user.Email = email
	users[id] = user
	return user, nil
}

// 🔹 Delete
func DeleteUser(id int) error {
	_, err := GetUser(id)
	if err != nil {
		return err
	}
	delete(users, id)
	return nil
}

// 🔹 Main function (test)
func main() {
	// Create
	u1 := CreateUser("Azizbek", "aziz@example.com")
	u2 := CreateUser("Ali", "ali@example.com")
	fmt.Println("Created Users:", u1, u2)

	// Get
	u, err := GetUser(1)
	if err == nil {
		fmt.Println("Get User 1:", u)
	}

	// Update
	updatedUser, err := UpdateUser(2, "Ali Updated", "ali_new@example.com")
	if err == nil {
		fmt.Println("Updated User 2:", updatedUser)
	}

	// Delete
	err = DeleteUser(1)
	if err == nil {
		fmt.Println("Deleted User 1")
	}

	// Try to get deleted user
	_, err = GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
