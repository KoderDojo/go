package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home Page")
	})

	http.HandleFunc("/landing", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "You've been redirected.")
	})

	// Routes to redirect.
	redirectedRoutes := make(map[string]string)
	redirectedRoutes["/redirect"] = "/landing"
	redirectedRoutes["/missing"] = "/landing"

	for k, v := range redirectedRoutes {
		http.Handle(k, http.RedirectHandler(v, http.StatusFound))
	}

	fmt.Println("Listening on port 8080...")

	http.ListenAndServe(":8080", nil)
}
