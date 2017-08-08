package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	p := struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Age       int    `json:"age"`
	}{
		"John",
		"Doe",
		15,
	}

	val, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(val)
	fmt.Println(string(val))
}
