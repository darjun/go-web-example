package main

import (
	"log"
	"math/rand"
	"os"
	"text/template"
	"time"
)

type AgeInfo struct {
	Age				int
	GreaterThan60 	bool
	GreaterThan40 	bool
}

func main() {
	t, err := template.ParseFiles("test")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	rand.Seed(time.Now().Unix())
	age := rand.Intn(100)
	info := AgeInfo {
		Age:			age,
		GreaterThan60: 	age > 60,
		GreaterThan40: 	age > 40,
	}
	err = t.Execute(os.Stdout, info)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}