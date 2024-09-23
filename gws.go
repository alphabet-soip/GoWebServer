package main

import (
	"github.com/alphabet-soip/GoWebServer"
	"fmt"
	"net/http"
)


func main() {
	// Route for the root path
	http.HandleFunc("/hello-world-html", api.helloWorldHTML)
	http.HandleFunc("/hello-world", api.helloWorld)
	http.HandleFunc("/hello-world-json", api.helloWorldJSON)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
