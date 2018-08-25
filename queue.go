package main

import "strconv"

type Queue struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Question     string `json:"question"`
	Googled      bool   `json:"googled"`
	AskedStudent bool   `json:"askedStudent"`
	HasDebugged  bool   `json:"hasDebugged"`
	Contacted    bool   `json:"contacted"`
	Completed    bool   `json:"completed"`
}

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
