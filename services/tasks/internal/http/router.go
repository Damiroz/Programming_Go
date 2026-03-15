package http

import (
    "net/http"

    "tech-ip-sem2/services/tasks/client/authclient"
    "tech-ip-sem2/services/tasks/internal/service"
    "tech-ip-sem2/shared/middleware"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    tasksService := service.NewTasksService()
    authClient := authclient.NewClient()
    h := NewHandler(tasksService, authClient)

    mux.HandleFunc("/v1/tasks", h.TasksCollection)
    mux.HandleFunc("/v1/tasks/", h.TaskItem)

    return middleware.RequestID(middleware.Logging(mux))
}
