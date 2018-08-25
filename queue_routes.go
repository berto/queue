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

	s.HandleFunc("/{id}", deleteQueue).Methods("DELETE")
}

func listQueues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := QueueResponse{
		Error: "",
		Data:  mockQueues(),
	}
	json.NewEncoder(w).Encode(response)
}

func addQueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := QueueResponse{
		Error: "",
		Data:  []Queue{mockQueue()},
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
		queue = mockQueue()
		queue.ID = id
		queue.Completed = true
	}
	response := QueueResponse{
		Error: errorMessage,
		Data:  []Queue{queue},
	}
	json.NewEncoder(w).Encode(response)
}
