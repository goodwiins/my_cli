package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Message represents a JSON message.
type Message struct {
	Text string `json:"text"`
}

// homeHandler handles the root path.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Welcome to my HTTP server!")
}

// apiHandler handles a simple API endpoint.
func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "GET request received")
	case http.MethodPost:
		fmt.Fprint(w, "POST request received")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := map[string]string{"received": msg.Text}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/json", jsonHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
