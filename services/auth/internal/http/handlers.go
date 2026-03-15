package http

import (
    "encoding/json"
    "net/http"
    "strings"

    "tech-ip-sem2/services/auth/internal/service"
)

type Handler struct {
    auth *service.AuthService
}

func NewHandler(auth *service.AuthService) *Handler {
    return &Handler{auth: auth}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req service.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    resp, err := h.auth.Login(req)
    if err != nil {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        writeVerifyResponse(w, service.VerifyResponse{
            Valid: false,
            Error: "unauthorized",
        }, http.StatusUnauthorized)
        return
    }

    parts := strings.SplitN(authHeader, " ", 2)
    if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
        writeVerifyResponse(w, service.VerifyResponse{
            Valid: false,
            Error: "unauthorized",
        }, http.StatusUnauthorized)
        return
    }

    token := parts[1]
    resp, _ := h.auth.VerifyToken(token)

    status := http.StatusOK
    if !resp.Valid {
        status = http.StatusUnauthorized
    }

    writeVerifyResponse(w, resp, status)
}

func writeVerifyResponse(w http.ResponseWriter, resp service.VerifyResponse, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(resp)
}
