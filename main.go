package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/boonyarit-iamsaard/templ-htmx-go/handlers"
)

func main() {
	handler := handlers.NewHandler()

	http.HandleFunc("/", handler.HomeHandler)

	// serve static files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
