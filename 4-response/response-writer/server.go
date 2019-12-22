package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
    FirstName   string      `json:"first_name"`
    LastName    string      `json:"last_name"`
    Age         int         `json:"age"`
    Hobbies     []string    `json:"hobbies"`
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web 编程之 响应</title></head>
<body><h1>直接使用 Write 方法<h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "This API not implemented!!!")
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://baidu.com")
	w.WriteHeader(302)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    u := &User {
        FirstName:  "lee",
        LastName:   "darjun",
        Age:        18,
        Hobbies:    []string{"coding", "math"},
    }
    data, _ := json.Marshal(u)
    w.Write(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/write", writeHandler)
	mux.HandleFunc("/writeheader", writeHeaderHandler)
	mux.HandleFunc("/header", headerHandler)
	mux.HandleFunc("/json", jsonHandler)
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
