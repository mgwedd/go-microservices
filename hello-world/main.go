package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mgwedd/go-microservices/basic/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	handleHello := handlers.NewHello(logger)
	handleGoodbye := handlers.NewGoodbye(logger)

	mux := http.NewServeMux()

	mux.Handle("/hello", handleHello)
	mux.Handle("/goodbye", handleGoodbye)

	server := &http.Server{
		Addr:         ":9091",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	logger.Printf("Received terminal, graceful shutdown signal: %s", sig)

	timeoutCxt, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutCxt)
}
