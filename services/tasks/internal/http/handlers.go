package http

import (
    "encoding/json"
    "errors"
    "log"
    "net/http"
    "strings"

    "tech-ip-sem2/services/tasks/client/authclient"
    "tech-ip-sem2/services/tasks/internal/service"
    "tech-ip-sem2/shared/middleware"
)

type Handler struct {
    tasks *service.TasksService
    auth  *authclient.Client
}

func NewHandler(tasks *service.TasksService, auth *authclient.Client) *Handler {
    return &Handler{
        tasks: tasks,
        auth:  auth,
    }
}

// /v1/tasks (POST, GET)
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        h.handleCreateTask(w, r)
    case http.MethodGet:
        h.handleListTasks(w, r)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

// /v1/tasks/{id} (GET, PATCH, DELETE)
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
    if id == "" {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }

    switch r.Method {
    case http.MethodGet:
        h.handleGetTask(w, r, id)
    case http.MethodPatch:
        h.handleUpdateTask(w, r, id)
    case http.MethodDelete:
        h.handleDeleteTask(w, r, id)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

func (h *Handler) authorize(w http.ResponseWriter, r *http.Request) bool {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return false
    }

    parts := strings.SplitN(authHeader, " ", 2)
    if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return false
    }

    token := parts[1]
    reqID := middleware.GetRequestID(r)

    vr, status, err := h.auth.Verify(r.Context(), token, reqID)
    if err != nil {
        log.Printf("auth verify error: %v", err)
        http.Error(w, "auth service unavailable", http.StatusBadGateway)
        return false
    }

    if status == http.StatusOK && vr.Valid {
        return true
    }

    if status == http.StatusUnauthorized {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return false
    }

    http.Error(w, "auth error", http.StatusBadGateway)
    return false
}

func (h *Handler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
    if !h.authorize(w, r) {
        return
    }

    var req service.CreateTaskRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    task, err := h.tasks.Create(req)
    if err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

func (h *Handler) handleListTasks(w http.ResponseWriter, r *http.Request) {
    if !h.authorize(w, r) {
        return
    }

    tasks := h.tasks.List()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) handleGetTask(w http.ResponseWriter, r *http.Request, id string) {
    if !h.authorize(w, r) {
        return
    }

    task, ok := h.tasks.Get(id)
    if !ok {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func (h *Handler) handleUpdateTask(w http.ResponseWriter, r *http.Request, id string) {
    if !h.authorize(w, r) {
        return
    }

    var req service.UpdateTaskRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    task, err := h.tasks.Update(id, req)
    if err != nil {
        if errors.Is(err, errors.New("not found")) {
            http.Error(w, "not found", http.StatusNotFound)
            return
        }
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func (h *Handler) handleDeleteTask(w http.ResponseWriter, r *http.Request, id string) {
    if !h.authorize(w, r) {
        return
    }

    if err := h.tasks.Delete(id); err != nil {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
