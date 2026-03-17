package http

import "net/http"

func NewRouter(h *Handler) http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/v1/tasks", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            h.CreateTask(w, r)
            return
        }
        if r.Method == http.MethodGet {
            h.GetTasks(w, r)
            return
        }
    })

    return mux
}
