package main

import (
	"log"
	"os"
	"text/template"
)

type Item struct {
	Name	string
	Price	float64
	Num		int
}

func (item Item) Total() float64 {
	return item.Price * float64(item.Num)
}

func main() {
	t, err := template.ParseFiles("test")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	item := Item {"iPhone", 5499.99, 2 }

	err = t.Execute(os.Stdout, item)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}