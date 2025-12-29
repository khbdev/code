package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type User struct {
	ID int 	`json:"id"`
	Name string `json:"name"`
}


var (
	users = []User{}
	nextId = 1
	my sync.Mutex
)


func GetUsers(w http.ResponseWriter, r *http.Request){
	my.Lock()
	defer my.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func PostUsers(w http.ResponseWriter, r *http.Request){
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid Create User", http.StatusBadRequest)
		return
	}

	my.Lock()
	u.ID = nextId
	nextId++
	users = append(users, u)
	my.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func UserHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		GetUsers(w, r)
	case http.MethodPost:
		PostUsers(w, r)
	}
}

func main(){
	http.HandleFunc("/users", UserHandler)

	http.ListenAndServe(":8081", nil)
}