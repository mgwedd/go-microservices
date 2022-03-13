package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)

	hello.logger.Printf("Received Data %s\n", data)

	if err != nil {
		http.Error(writer, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "Hello, echoing your data: %s\n", data)
}
