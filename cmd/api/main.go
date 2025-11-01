package main

import (
	"fmt"
	"net/http"

	"shortest-path-api/internal/application"
	"shortest-path-api/internal/infrastructure/handlers"
)

func main() {
	service := application.NewShortestPathService()
	handler := handlers.NewHTTPHandler(service)

	//http.HandleFunc("/shortest-path", handler.ShortestPath)

	http.HandleFunc("/shortest-path", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		handler.ShortestPath(w, r)
	})

	fmt.Println("Servidor ejecutándose en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
