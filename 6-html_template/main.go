package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("hello.html"))
    t.Execute(w, "Hello World")
}

func conditionHandler(w http.ResponseWriter, r *http.Request) {
	age, err := strconv.ParseInt(r.URL.Query().Get("age"), 10, 64)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	t := template.Must(template.ParseFiles("condition.html"))
	t.Execute(w, age)
}

type Item struct {
	Name	string
	Price	int
}

func iterateHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("iterate.html"))

	items := []Item {
		{ "iPhone", 5499 },
		{ "iPad", 6331 },
		{ "iWatch", 1499 },
		{ "MacBook", 8250 },
	}
	t.Execute(w, items)
}

type User struct {
	Name	string
	Age		int
}

type Pet struct {
	Name	string
	Age		int
	Owner	User
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("set.html"))

	pet := Pet {
		Name:	"Orange",
		Age:	2,
		Owner:	User {
			Name:	"dj",
			Age:	28,
		},
	}
	t.Execute(w, pet)
}

func includeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("include1.html", "include2.html"))
	t.Execute(w, "Hello World!")
}

func pipelineHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("pipeline.html"))
	t.Execute(w, rand.Float64())
}

func formateDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func funcsHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{ "fdate": formateDate }
	t := template.Must(template.New("funcs.html").Funcs(funcMap).ParseFiles("funcs.html"))
	t.Execute(w, time.Now())
}

func contextAwareHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("context-aware.html"))
	t.Execute(w, `He saied: <i>"She's alone?"</i>`)
}

func xssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
	    // w.Header().Set("X-XSS-Protection", "0")
		t := template.Must(template.ParseFiles("xss-display.html"))
		t.Execute(w, template.HTML(r.FormValue("comment")))
	} else {
		t := template.Must(template.ParseFiles("xss-form.html"))
		t.Execute(w, nil)
	}
}

type NestInfo struct {
	Name string
	Todos []string
}

func nestHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("layout.html"))

	data := NestInfo {
		"dj", []string{"Homework", "Game", "Cleaning"},
	}
	t.ExecuteTemplate(w, "layout", data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/condition", conditionHandler)
	mux.HandleFunc("/iterate", iterateHandler)
	mux.HandleFunc("/set", setHandler)
	mux.HandleFunc("/include", includeHandler)
	mux.HandleFunc("/pipeline", pipelineHandler)
	mux.HandleFunc("/funcs", funcsHandler)
	mux.HandleFunc("/contextAware", contextAwareHandler)
	mux.HandleFunc("/xss", xssHandler)
	mux.HandleFunc("/nest", nestHandler)

	server := &http.Server {
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}