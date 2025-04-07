package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World from Azure App Service!")
}

func main() {
	http.HandleFunc("/hello", handler)

	// Get port from environment variable (required for Azure)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	fmt.Println("Server is running on port", port)
	http.ListenAndServe(":"+port, nil)
}
