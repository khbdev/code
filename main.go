package main

import "fmt"

func main() {
	users := []string{}

	for i := 0; i < 5; i++ {
		users = append(users, "Azizbek")
		users = append(users, "Soliha")
		users = append(users, "Ali")
		users = append(users, "Muslima")
	}

	fmt.Println("Users", users)
	fmt.Println("uzunlik", len(users))

}
