package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.RawQuery)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	server := &http.Server {
		Addr:		":8080",
		Handler:	mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}