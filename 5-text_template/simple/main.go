package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

func stringLiteralTemplate() {
	s := "My name is {{ .Name }}. I am {{ .Age }} years old.\n"
	t, err := template.New("test").Parse(s)
	if err != nil {
		log.Fatal("Parse string literal template error:", err)
	}

	u := User{Name: "darjun", Age: 28}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute string literal template error:", err)
	}
}

func fileTemplate() {
	t, err := template.ParseFiles("test")
	if err != nil {
		log.Fatal("Parse file template error:", err)
	}

	u := User{Name: "dj", Age: 18}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute file template error:", err)
	}
}

func main() {
	stringLiteralTemplate()

	fileTemplate()
}
