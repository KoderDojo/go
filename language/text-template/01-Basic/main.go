package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	name := "John Doe"

	tmpl, err := template.New("template").Parse("Hi, my name is {{.}}.")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(os.Stdout, name)
}
