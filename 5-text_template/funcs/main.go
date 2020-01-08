package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func formatDate(t time.Time) string {
	return t.Format("2016-01-02")
}

func main() {
	funcMap := template.FuncMap {
		"fdate": formatDate,
	}
	t := template.New("test").Funcs(funcMap)
	t, err := t.ParseFiles("test")
	if err != nil {
		log.Fatal("Parse errr:", err)
	}

	err = t.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatal("Exeute error:", err)
	}
}
