package main

import (
	"fmt"
	"log"
	// "io/ioutil"
	"net/http"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	URL := r.URL

	fmt.Fprintf(w, "Scheme: %s\n", URL.Scheme)
	fmt.Fprintf(w, "Host: %s\n", URL.Host)
	fmt.Fprintf(w, "Path: %s\n", URL.Path)
	fmt.Fprintf(w, "RawPath: %s\n", URL.RawPath)
	fmt.Fprintf(w, "RawQuery: %s\n", URL.RawQuery)
	fmt.Fprintf(w, "Fragment: %s\n", URL.Fragment)
}

func protoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
    fmt.Fprintf(w, "ProtoMajor: %d\n", r.ProtoMajor)
    fmt.Fprintf(w, "ProtoMinor: %d\n", r.ProtoMinor)
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
    for key, value := range r.Header {
        fmt.Fprintf(w, "%s: %v\n", key, value)
    }
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
    data := make([]byte, r.ContentLength)
	r.Body.Read(data) // 忽略错误处理
	// data := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    
    fmt.Fprintln(w, string(data))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `
<html>
    <head>
        <title>Go Web 编程之 request</title>
    </head>
    <body>
        <form method="post" action="/body">
            <label for="username">用户名：</label>
            <input type="text" id="username" name="username">
            <label for="email">邮箱：</label>
            <input type="text" id="email" name="email">
            <button type="submit">提交</button>
        </form>
    </body>
</html>
`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/url", urlHandler)
	mux.HandleFunc("/proto", protoHandler)
	mux.HandleFunc("/header", headerHandler)
	mux.HandleFunc("/body", bodyHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
