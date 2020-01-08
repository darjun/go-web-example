package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("test1", "test2")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	err = t.Execute(os.Stdout, "test data")
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}