package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	port := os.Getenv("PORT")

	r := createRouter()

	log.Printf("server running on port: %v", port)
	http.ListenAndServe(":"+port, r)
}

func createRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", pong)
	applyQueueRoutes(r)

	return r
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
