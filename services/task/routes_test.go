package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mariosker/taskfrenzy/types"
)

func TestTaskServiceHandlers(t *testing.T) {
	taskStore := &mockTaskStore{}
	userStore := &mockUserStore{}
	handler := NewHandler(taskStore, userStore)

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
type mockUserStore struct{}

func (m *mockTaskStore) CreateTask(task types.Task) error { return nil }
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) CreateUser(user types.User) error { return nil }

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) { return nil, nil }
