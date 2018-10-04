package main

import (
	"strconv"
	"time"
)

// Queue represents a queue question
type Queue struct {
	ID           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Location     string    `db:"location" json:"location"`
	Question     string    `db:"question" json:"question"`
	Googled      bool      `db:"googled" json:"googled"`
	AskedStudent bool      `db:"asked_student" json:"askedStudent"`
	HasDebugged  bool      `db:"has_debugged" json:"hasDebugged"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
	CompletedAt  time.Time `db:"completed_at" json:"completedAt"`
	Contacted    bool      `db:"contacted" json:"contacted"`
	Completed    bool      `db:"completed" json:"completed"`
}

// QueueResponse represents the HTTP body response
type QueueResponse struct {
	Error string  `json:"error"`
	Data  []Queue `json:"data"`
}

func mockQueue() Queue {
	return Queue{
		ID:       1,
		Name:     "Bob",
		Location: "Classroom 2",
		Question: "What's a git?",
		Googled:  true,
	}
}

func mockQueues() []Queue {
	names := []string{"Bob", "Cat", "Tow"}
	questions := []string{"What's a git?", "I can't get my route to post", "Halp!"}
	length := len(names)
	queues := make([]Queue, length)
	for i := 0; i < length; i++ {
		queues[i] = Queue{
			ID:       i,
			Name:     names[i],
			Location: "Classroom" + strconv.Itoa(i),
			Question: questions[i],
			Googled:  true,
		}
	}
	return queues
}
