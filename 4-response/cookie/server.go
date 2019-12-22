package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := &http.Cookie{
		Name:     "name",
		Value:    "darjun",
		HttpOnly: true,
	}
	c2 := &http.Cookie{
		Name:     "age",
		Value:    "18",
		HttpOnly: true,
	}
	w.Header().Set("set-cookie", c1.String())
	w.Header().Add("set-cookie", c2.String())
}

func setCookie2(w http.ResponseWriter, r *http.Request) {
	c1 := &http.Cookie{
		Name:     "name",
		Value:    "darjun",
		HttpOnly: true,
	}
	c2 := &http.Cookie{
		Name:     "age",
		Value:    "18",
		HttpOnly: true,
	}
	http.SetCookie(w, c1)
	http.SetCookie(w, c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Host:", r.Host)
	fmt.Fprintln(w, "Cookies:", r.Header["Cookie"])
}

func getCookie2(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("name")
	if err != nil {
		fmt.Fprintln(w, "cannot get cookie of name")
	}

	cookies := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cookies)
}

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/set_cookie", setCookie)
	mux1.HandleFunc("/set_cookie2", setCookie2)
	mux1.HandleFunc("/get_cookie", getCookie)
	mux1.HandleFunc("/get_cookie2", getCookie2)

	server1 := &http.Server{
		Addr:    ":8080",
		Handler: mux1,
	}

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/get_cookie", getCookie)

	server2 := &http.Server {
		Addr: 	 ":8081",
		Handler: mux2,
	}
	
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func () {
		defer wg.Done()

		if err := server1.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	
	go func() {
		defer wg.Done()

		if err := server2.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
