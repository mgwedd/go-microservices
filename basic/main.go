package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mgwedd/go-microservices/basic/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	helloHandler := handlers.NewHello(logger)
	goodbyeHandler := handlers.NewGoodbye(logger)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	http.ListenAndServe(":9091", serveMux)
}
