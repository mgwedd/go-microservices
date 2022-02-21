package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mgwedd/go-microservices/basic/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	handleHello := handlers.NewHello(logger)
	handleGoodbye := handlers.NewGoodbye(logger)

	mux := http.NewServeMux()

	mux.Handle("/", handleHello)
	mux.Handle("/goodbye", handleGoodbye)

	server := &http.Server{
		Addr:         ":9091",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()
}
