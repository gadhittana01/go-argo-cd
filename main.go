package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status": "okay giri putra adhittana"}`)
}

func main() {
	http.HandleFunc("/health", healthCheck)

	fmt.Println("Server starting on port 8080...")
	fmt.Println("Health check available at: http://localhost:8080/health")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
