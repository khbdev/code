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


	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Salom")
	})

	mux.HandleFunc("/salom", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	log.Println("request keldi")

	time.Sleep(8 * time.Second)

	log.Println("request tugadi")

	fmt.Fprint(w, "Request tugadi")
})


	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)


	server := &http.Server{
		Addr: ":8084",
		Handler: mux,
	}

	go func() {
		fmt.Println("Server ishga tushdi Port :8084")
		server.ListenAndServe()
	}()

	sig := <-sigCh
	fmt.Println("Signal keldi: ", sig)

	fmt.Println("ShutDown ishga tushdi 10s")

	ctx, cance := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cance()

	server.Shutdown(ctx)
	server.Close()

	fmt.Println("Server ochdi !")
}

