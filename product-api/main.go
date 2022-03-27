package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/mgwedd/go-microservices/product-api/handlers"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	handleProducts := handlers.NewProducts(logger)

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/products", handleProducts.GetProducts)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/products", handleProducts.AddProduct)
	postRouter.Use(handleProducts.ValidateRequest)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/api/products/{id:[0-9]+}", handleProducts.UpdateProduct)
	putRouter.Use(handleProducts.ValidateRequest)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/api/products/{id:[0-9]+}", handleProducts.DeleteProduct)

	server := &http.Server{
		Addr:         *bindAddress,
		Handler:      router,
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
