package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	var p Person

	err := json.Unmarshal([]byte(`{"firstname":"John","lastname":"Doe","age":15}`), &p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p)
}
