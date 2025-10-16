// File: frontend/frontend.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Structure for API requests (e.g., for submitting articles)
type Article struct {
	Title  string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
}

// Handler for submitting an article (proxies to Python backend)
func submitArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Forward the request to the Python backend
	resp, err := http.Post("http://localhost:5000/submit-article", "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error connecting to backend", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response from Python backend
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading backend response", http.StatusInternalServerError)
		return
	}

	// Write the response back to the client
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)
}

// Handler for getting articles (proxies to Python backend)
func getArticlesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:5000/articles")
	if err != nil {
		http.Error(w, "Error connecting to backend", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading backend response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)
}

// Main function to set up the server
func main() {
	http.HandleFunc("/submit", submitArticleHandler)    // Endpoint: POST http://localhost:8080/submit
	http.HandleFunc("/articles", getArticlesHandler)   // Endpoint: GET http://localhost:8080/articles

	fmt.Println("OSINT Journal Frontend running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
