package main

import (
	"fmt"
	"net/http"
)

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From %s", r.URL.Path)
}

func main() {
	var s server

	fmt.Println("Listening...")

	http.ListenAndServe(":8080", s)
}
