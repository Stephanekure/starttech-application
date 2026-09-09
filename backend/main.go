package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type LogMessage struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Redis   string `json:"redis_host,omitempty"`
}

func logJSON(level, message string) {
	entry := LogMessage{
		Level:   level,
		Message: message,
		Redis:   os.Getenv("REDIS_HOST"),
	}
	jsonBytes, _ := json.Marshal(entry)
	log.Println(string(jsonBytes))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"healthy"}`))
}

func main() {
	// Read required environment variables
	redisHost := os.Getenv("REDIS_HOST")
	mongoURI := os.Getenv("MONGO_URI")

	if redisHost == "" || mongoURI == "" {
		logJSON("WARN", "Missing REDIS_HOST or MONGO_URI environment variables")
	}

	http.HandleFunc("/api/v1/health", healthHandler)
	http.HandleFunc("/health", healthHandler)

	logJSON("INFO", "Starting Golang backend server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}