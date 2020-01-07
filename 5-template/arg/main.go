package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	FirstName 	string
	LastName	string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	t, err := template.ParseFiles("test")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	err = t.Execute(os.Stdout, User{FirstName: "lee", LastName: "darjun"})
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}