package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	var p Person

	rdr := strings.NewReader(`{"firstname":"John","lastname":"Doe","age":15}`)

	err := json.NewDecoder(rdr).Decode(&p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p)
}
