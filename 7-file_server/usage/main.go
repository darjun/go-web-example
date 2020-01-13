package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	ServeDir string
)

func init() {
	flag.StringVar(&ServeDir, "sd", "./", "the directory to serve")
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(ServeDir))))


	server := &http.Server {
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}