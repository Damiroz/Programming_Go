package http

import (
	"encoding/json"
	"net/http"
	"time"

	"tech-ip-sem2/services/tasks/internal/grpcclient"
	"tech-ip-sem2/services/tasks/internal/service"
)

type Handler struct {
	auth  *grpcclient.AuthClient
	tasks *service.TaskService
}

func NewHandler(auth *grpcclient.AuthClient, tasks *service.TaskService) *Handler {
	return &Handler{auth: auth, tasks: tasks}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	valid, _, err := h.auth.Verify(r.Context(), token)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusServiceUnavailable)
		return
	}
	if !valid {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var t service.Task
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = time.Now().Format("20060102150405")

	created := h.tasks.Create(t)
	json.NewEncoder(w).Encode(created)
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	valid, _, err := h.auth.Verify(r.Context(), token)
	if err != nil {
		http.Error(w, "auth unavailable", http.StatusServiceUnavailable)
		return
	}
	if !valid {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(h.tasks.GetAll())
}
