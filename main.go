package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userDetailHandler) // path param uchun
	http.HandleFunc("/search", searchHandler)

	fmt.Println("🚀 Server ishga tushdi: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// 1️⃣ Bosh sahifa
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h2>Salom, Azizbek!</h2><p>Go HTTP server ishlayapti ✅</p>")
}

// 2️⃣ Foydalanuvchilar ro‘yxati (static)
func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: "1", Name: "Azizbek"},
		{ID: "2", Name: "Coder"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// 3️⃣ Path param: /users/{id}
func userDetailHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // masalan /users/2
	parts := strings.Split(path, "/")

	if len(parts) < 3 || parts[2] == "" {
		http.NotFound(w, r)
		return
	}

	id := parts[2]
	user := User{ID: id, Name: "User #" + id}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// 4️⃣ Query param: /search?keyword=go
func searchHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		keyword = "bo‘sh"
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "🔍 Qidiruv natijasi: %s", keyword)
}
