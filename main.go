package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news/logger"
	"news/scraper"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func home(w http.ResponseWriter, r *http.Request) {
	// Only allow GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := ResponseMessage{
		Message: "Welcome to the NBC News API. Made by erxcc.",
	}
	// Return response as JSON
	json.NewEncoder(w).Encode(resp)
}

func nbcNews(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Gets the articles from NBC News
	articles, err := scraper.ScrapeNBC()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error getting news data."})
		return
	}
	// Return response as JSON
	json.NewEncoder(w).Encode(articles)
}

// Required for Vercel deployment
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		home(w, r)
	case "/nbc-news":
		nbcNews(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	// Routes
	http.HandleFunc("/", home)
	http.HandleFunc("/nbc-news", nbcNews)

	port := ":8080"
	logger.Yellow(fmt.Sprintf("Starting server on port %s", port))
	if err := http.ListenAndServe(port, nil); err != nil {
		logger.Red(fmt.Sprintf("Server failed to start: %v", err))
	}
}
