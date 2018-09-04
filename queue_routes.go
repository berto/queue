package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func applyQueueRoutes(r *mux.Router) {
	s := r.PathPrefix("/queue").Subrouter()

	s.HandleFunc("/", listQueues).Methods("GET")
	r.HandleFunc("/queue", listQueues).Methods("GET")

	s.HandleFunc("/", addQueue).Methods("POST")
	r.HandleFunc("/queue", addQueue).Methods("POST")

	s.HandleFunc("/{id}", updateQueue).Methods("PATCH")

	s.HandleFunc("/{id}", deleteQueue).Methods("DELETE")
}

func listQueues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	queues, err := getQueues()
	response := QueueResponse{
		Error: err,
		Data:  queues,
	}
	json.NewEncoder(w).Encode(response)
}

func addQueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	queue, err := parseBody(r)
	if err == "" {
		queue, err = insertQueue(queue)
	}
	response := QueueResponse{
		Error: err,
		Data:  []Queue{queue},
	}
	json.NewEncoder(w).Encode(response)
}

func updateQueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var errorMessage string
	var queue Queue
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorMessage = "Invalid ID"
	} else {
		queue, errorMessage = contactQueue(id)
	}
	response := QueueResponse{
		Error: errorMessage,
		Data:  []Queue{queue},
	}
	json.NewEncoder(w).Encode(response)
}

func deleteQueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var errorMessage string
	var queue Queue
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorMessage = "Invalid ID"
	} else {
		queue, errorMessage = completeQueue(id)
	}
	response := QueueResponse{
		Error: errorMessage,
		Data:  []Queue{queue},
	}
	json.NewEncoder(w).Encode(response)
}

func parseBody(r *http.Request) (Queue, string) {
	var queue Queue
	if r.Body == nil {
		return queue, "Ivalid request body"
	} else {
		err := json.NewDecoder(r.Body).Decode(&queue)
		if err != nil {
			return queue, "Ivalid queue json"
		} else {
			return queue, ""
		}
	}

}
