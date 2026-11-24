package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"` // would not be send to response if not capital letter
	Done  bool   `json:"done"`  // struct tags
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos := []Todo{{ID: 1, Title: "Learn Go", Done: false}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	})

	log.Println("listening on port 8008")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
