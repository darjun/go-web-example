package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func emptyMainTemplate() {
	t := template.New("test")
	t, err := t.ParseFiles("test1")

	if err != nil {
		log.Fatal("in emptyMainTemplate parse error:", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("in emptyMainTemplate execute error:", err)
	}
}

func associatedTemplate() {
	t := template.New("test")
	t, err := t.ParseFiles("test1")

	if err != nil {
		log.Fatal("in associatedTemplate parse error:", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "test1", nil)
	if err != nil {
		log.Fatal("in associatedTemplate execute error:", err)
	}
}

func globTemplate() {
	t, err := template.ParseGlob("tmpl*.glob")
	if err != nil {
		log.Fatal("in globTemplate parse error:", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 3; i++ {
		err = t.ExecuteTemplate(os.Stdout, fmt.Sprintf("tmpl%d.glob", i), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	// emptyMainTemplate()

	// associatedTemplate()

	globTemplate()
}