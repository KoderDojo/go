package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// TimeoutHandler times out after 5 seconds, because
	// long running process takes 15 seconds.
	http.Handle("/", http.TimeoutHandler(longRunningProcess(), 5*time.Second, "Taking Too Long!"))

	fmt.Println("Listening on port 8080...")

	http.ListenAndServe(":8080", nil)
}

// A long running process that takes 15 seconds.
func longRunningProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(15 * time.Second)
	}
}
