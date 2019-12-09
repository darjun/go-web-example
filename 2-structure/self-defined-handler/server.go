package main

import (
    "fmt"
    "log"
    "net/http"
)

type GreetingHandler struct {
    Language string
}

func (h GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", h.Language)
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/chinese", GreetingHandler{Language: "你好"})
    mux.Handle("/english", GreetingHandler{Language: "Hello"})
    
    server := &http.Server {
        Addr:   ":8080",
        Handler: mux,
    }
    
    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
