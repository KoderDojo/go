package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	data := []string{"John Doe", "Jane Doe", "Timmy Doe"}

	tmpl, err := template.New("template").Parse("{{ range .}} - {{.}}{{end}}")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(os.Stdout, data)
}
