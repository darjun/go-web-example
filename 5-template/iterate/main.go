package main

import (
	"log"
	"os"
	"text/template"
)

type Item struct {
	Name	string
	Price	int
}

func main() {
	t, err := template.ParseFiles("test")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	// items := []Item {
	// 	{ "iPhone", 5499 },
	// 	{ "iPad", 6331 },
	// 	{ "iWatch", 1499 },
	// 	{ "MacBook", 8250 },
	// }
	var items []Item

	err = t.Execute(os.Stdout, items)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}