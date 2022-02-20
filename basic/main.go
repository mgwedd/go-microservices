package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadAll(req.Body)

		log.Printf("Received Data %s\n", data)

		if err != nil {
			http.Error(writer, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(writer, "Hello, echoing your data: %s\n", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye, world")
	})
	http.ListenAndServe(":9091", nil)
}
