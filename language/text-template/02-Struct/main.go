package main

import (
	"log"
	"os"
	"text/template"
)

type data struct {
	Name string
	Age  int
}

func main() {
	d := data{"John Doe", 15}

	tmpl, err := template.New("template").Parse("Hi, my name is {{.Name}}. I am {{.Age}} years old.")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(os.Stdout, d)
}
