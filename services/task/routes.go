package task

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/mariosker/taskfrenzy/services/auth"
	"github.com/mariosker/taskfrenzy/types"
	"github.com/mariosker/taskfrenzy/utils"
)

type Handler struct {
	store     types.TaskStore
	userStore types.UserStore
}

func NewHandler(store types.TaskStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", auth.WithJWTAuth(h.createTask, h.userStore)).Methods("POST")
}

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var task types.CreateTaskPayload

	if err := utils.ParseJSON(r, &task); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(task); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateTask(types.Task{
		Title:       task.Title,
		Description: task.Description,
		UserId:      task.UserId,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusCreated, nil)

}
