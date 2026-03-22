package main

import (
	"fmt"
	"net/http"
	"time"
)



func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return  func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        method := r.Method
        path := r.URL.Path

        next(w, r)

        latency := time.Since(start)

        fmt.Printf("Mathod %d, Path %d, Latency %d", method, path, latency)
    }
}

func ResoverMiddleware(next http.HandlerFunc) http.HandlerFunc  {
    return  func(w http.ResponseWriter, r *http.Request) {
        defer func ()  {
            if err := recover(); err != nil {
                fmt.Println("Panic ushlab qolindi")
                http.Error(w, "server error", http.StatusInternalServerError)
            }
        }()
        next(w, r)
    }
}

func  AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return  func(w http.ResponseWriter, r *http.Request) {
        key := r.RemoteAddr

        if key != "secret123" {
            http.Error(w, "Xatolik", http.StatusBadRequest)
            return 
        }
        next(w, r)
    }
}
   





func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":8080", nil)
}