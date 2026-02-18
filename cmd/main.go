package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}


var (
	users []User
	lastId int
)

func writeJSOn(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSOn(w, status, map[string]string{"error": msg})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}

	var req CreateUserRequest 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
	}

	lastId++
	user := User{
		ID: lastId,
		Name: req.Name,
	}
	users = append(users, user)
	

	writeJSOn(w, http.StatusCreated, user)
}


func GetUserHadler(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}

	writeJSOn(w, http.StatusOK, users)
}


func SwitchHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		GetUserHadler(w, r)
	case "POST":
		CreateUserHandler(w, r)
	}
}

func main(){
	http.HandleFunc("/users", SwitchHandler)
	fmt.Println("Server running :8086")
	http.ListenAndServe(":8086", nil)
}