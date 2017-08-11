package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	p := Person{"John", "Doe", 15}
	msg, _ := json.MarshalIndent(p, "", "    ")
	fmt.Println(string(msg))

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(p)
}
