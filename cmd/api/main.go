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

	http.HandleFunc("/shortest-path", handler.ShortestPath)

	fmt.Println("Servidor ejecutándose en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
