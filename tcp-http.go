package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	// Create tcp listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error creating listener: ", err)
	}

	// Create HTTP server that handles incoming request
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>Http on TCP<h1>")
	})

	// Start serving request using listener
	log.Println("Starting HTTP server on port 8080")
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving requests: ", err)
	}

}
