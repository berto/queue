package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestQueueRoutes(t *testing.T) {
	dbName := os.Getenv("DB_TEST_NAME")
	os.Setenv("DB_NAME", dbName)

	err := cleanDB()
	if err != "" {
		t.Errorf("Failed to clean db: %s.", err)
	}

	r := createRouter()

	tt := []struct {
		name    string
		method  string
		uri     string
		payload []byte
	}{
		{"get queues", "GET", "/queue", nil},
		{"post new queue", "POST", "/queue", []byte(`{"name": "test"}`)},
		{"update queue contacted status", "PATCH", "/queue/1", nil},
		{"delete queue by marking as complete", "DELETE", "/queue/1", nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var payload io.Reader
			if tc.payload != nil {
				payload = bytes.NewBuffer(tc.payload)
			}
			req, err := http.NewRequest(tc.method, tc.uri, payload)
			if err != nil {
				t.Errorf("Get failed with error %d.", err)
			}

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			var body QueueResponse
			json.NewDecoder(resp.Body).Decode(&body)

			if resp.Code != 200 {
				t.Errorf("/queue failed with error code %d.", resp.Code)
			}

			if resp.Header().Get("Content-Type") != "application/json; charset=utf-8" {
				t.Errorf("/queue failed with incorrect json header: %v", resp.Header().Get("Content-Type"))
			}

			if body.Error != "" && isQueueList(body.Data) {
				t.Errorf("/queue failed with incorrect response data")
			}
		})
	}
}

func isQueueList(data interface{}) bool {
	switch data.(type) {
	case []Queue:
		return true
	default:
		return false
	}
}
