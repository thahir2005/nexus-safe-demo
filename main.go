package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"healthy","service":"nexus-safe-demo"}`)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", handler)

	fmt.Println("NEXUS safe demo listening on :8000")
	http.ListenAndServe(":8000", nil)
}
