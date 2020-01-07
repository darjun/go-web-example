package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name	string
	Age		int
}

type Pet struct {
	Name	string
	Age		int
	Owner	User
}

func main() {
	t, err := template.ParseFiles("test")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	p := Pet {
		Name:	"Orange",
		Age:	2,
		Owner:	User {
			Name:	"dj",
			Age:	28,
		},
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}