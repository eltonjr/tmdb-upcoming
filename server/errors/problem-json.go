package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ProblemJSON will be returned by the server in case of an error.
// It follows the schema defined by RFC7807 for returning errors
type ProblemJSON struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Status int    `json:"status,omitempty"`
}

// Respond will write a ProblemJSON to responseWriter.
// Errors inside this function will be suppressed
func Respond(w http.ResponseWriter, status int, _type, title string) {
	json, _ := json.Marshal(ProblemJSON{
		Type:   _type,
		Title:  title,
		Status: status,
	})

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "%s", json)
}
