package main

import (
	"fmt"
	"net/http"
	"time"
)



func LoggerMiddleware(next  http.HandlerFunc) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {
		start := time.Now

		method := r.Method
		path := r.URL.Path

		next(w, r)

		latency := time.Since(start())

		fmt.Printf("Mathod %d, Path %d, Latency %d", method, path, latency)
	}
}


func PanicMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {
		defer func ()  {
			if err := recover(); err != nil {
				fmt.Println("Panic ushlab qolidi")
				http.Error(w, "Error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {
		key := r.Form.Get("X-Api-Key")
		if key != "secret123" {
			http.Error(w, "Key yoq", http.StatusUnauthorized)
			return 
		}
		next(w, r)
	}
}