package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	p := Person{"John", "Doe", 15}

	// Encode to Stdout
	err := json.NewEncoder(os.Stdout).Encode(p)
	if err != nil {
		log.Fatalln(err)
	}

	// Encode to buffer and print as string
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(b.String())
}
