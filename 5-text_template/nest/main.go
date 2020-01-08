package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("layout.tmpl")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "layout", "amazing")
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}