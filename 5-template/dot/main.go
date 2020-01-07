package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("(name:%s age:%d)", u.Name, u.Age)
}

func main() {
	s := "The user is {{ . }}."
	t, err := template.New("test").Parse(s)
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	u := User{Name: "darjun", Age: 28}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}