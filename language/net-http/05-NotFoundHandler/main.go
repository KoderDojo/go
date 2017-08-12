package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Page")
	})

	// Routes to return HTTP 404 Status.
	removedRoutes := make(map[string]struct{})
	var notFound struct{}
	removedRoutes["/removed"] = notFound
	removedRoutes["/removed2"] = notFound

	for k := range removedRoutes {
		http.Handle(k, http.NotFoundHandler())
	}

	fmt.Println("Listening on port 8080...")

	http.ListenAndServe(":8080", nil)
}
