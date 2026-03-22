package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)



func main(){
mux := http.NewServeMux()

mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println(w, "Salom")
})

mux.HandleFunc("/hoyy", func(w http.ResponseWriter, r *http.Request) {
  log.Printf("Request keldi")
  time.Sleep(11 * time.Second)
   log.Printf("Request tugadi")
   fmt.Println(w, "Request tugadi")
})


sigCh := make(chan os.Signal, 1)

signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

server := &http.Server{
    Addr: ":8083",
    Handler: mux,
}


go func() {
    fmt.Println("Server Running :8083")
    server.ListenAndServe()
}()


sig := <-sigCh
fmt.Println("Signal keldi", sig)

ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
defer cancel()

server.Shutdown(ctx)
server.ListenAndServe()

fmt.Println("Server ochdi")
}