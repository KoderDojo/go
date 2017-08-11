package main

import (
	"fmt"
	"net/http"
)

// Wassup is an informal greeter that implements Handler interface.
type Wassup struct{}

func (Wassup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Wassup!")
}

func main() {
	var wassup Wassup

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello From %s", r.URL.Path)
	})

	http.Handle("/greet/", wassup)

	fmt.Println("Listening on port 8080...")

	http.ListenAndServe(":8080", nil)
}
