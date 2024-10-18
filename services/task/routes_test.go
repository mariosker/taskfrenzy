package task

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mariosker/taskfrenzy/types"
)

func TestTaskServiceHandlers(t *testing.T) {
	taskStore := &mockTaskStore{}
	handler := NewHandler(taskStore)

	t.Run("should fail if the task payload is invalid", func(t *testing.T) {
		payload := types.CreateTaskPayload{
			Description: "hhh",
			Title:       "",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", handler.createTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should work correctly", func(t *testing.T) {
		payload := types.CreateTaskPayload{
			Description: "hhh",
			Title:       "asadss",
			UserId:      "2",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", handler.createTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockTaskStore struct{}

func (m *mockTaskStore) CreateTask(task types.Task) error { return nil }
