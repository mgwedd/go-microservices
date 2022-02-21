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
	serveMux := http.NewServeMux()

	serveMux.Handle("/", helloHandler)

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		logger.Println("goodbye, world")
	})

	http.ListenAndServe(":9091", serveMux)
}
