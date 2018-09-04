package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	port := os.Getenv("PORT")

	r := createRouter()

	handler := applyCorsMiddleware(r)
	log.Printf("server running on port: %v", port)
	http.ListenAndServe(":"+port, handler)
}

func createRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", pong)
	applyQueueRoutes(r)

	return r
}

func applyCorsMiddleware(r *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	})
	handler := c.Handler(r)
	return handler
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
