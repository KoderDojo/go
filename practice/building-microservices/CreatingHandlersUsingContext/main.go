// curl localhost:8080/helloWorld -d '{"name":"John"}'

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type validationContextKey string

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

type validationHandler struct {
	next http.Handler
}

func (h validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(w, r)
}

func newValidationHandler(next http.Handler) validationHandler {
	return validationHandler{next}
}

func main() {
	handler := newValidationHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Context().Value(validationContextKey("name")).(string)
		response := helloWorldResponse{"hi " + name}

		encoder := json.NewEncoder(w)
		encoder.Encode(response)
	}))

	http.Handle("/helloWorld", handler)

	fmt.Println("Listening on port 8080...")

	http.ListenAndServe(":8080", nil)
}
