package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	data := make(map[string]interface{})
	data["Name"] = "John Doe"
	data["Age"] = 15
	data["Address"] = struct {
		Line1 string
		City  string
		State string
		Zip   string
	}{
		"1234 Broadway Street",
		"Anywhere",
		"UT",
		"80323",
	}

	tmpl, err := template.New("template").Parse(
		`Hi, my name is {{.Name}}.
		I am {{.Age}} years old.
		I live at {{.Address.Line1}} {{.Address.City}}, {{.Address.State}}  {{.Address.Zip}}.`)

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(os.Stdout, data)
}
