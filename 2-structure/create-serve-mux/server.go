package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
	}

	i err := server.ListenAndServe(); err != nil {
	log.Fatal(err)
	}
}
