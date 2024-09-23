package api

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func helloWorldHTML(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type to HTML
	w.Header().Set("Content-Type", "")
	// Write the HTML response
	fmt.Fprintf(w, "<h1>Hello World! - GWS</h1>")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Text response
	fmt.Fprintf(w, "Hello World! - GWS")
}

func helloWorldJSON(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello World! - GWS"}
	json.NewEncoder(w).Encode(response)
}
