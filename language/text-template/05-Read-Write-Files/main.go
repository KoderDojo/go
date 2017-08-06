package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	data := struct {
		Name    string
		Hobbies []string
	}{
		"John Doe",
		[]string{"baseball", "basketball", "football", "tennis"},
	}

	tmpl, err := template.ParseFiles("template.gotmpl")
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	tmpl.Execute(out, data)
}
