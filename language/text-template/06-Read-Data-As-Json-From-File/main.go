package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

func main() {
	var data struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Hobbies []string `json:"hobbies"`
	}

	dataFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(dataFile)
	if err = jsonParser.Decode(&data); err != nil {
		log.Fatal(err)
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
