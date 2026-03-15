package http

import (
    "net/http"

    "tech-ip-sem2/services/auth/internal/service"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    authService := service.NewAuthService()
    h := NewHandler(authService)

    mux.HandleFunc("/v1/auth/login", h.Login)
    mux.HandleFunc("/v1/auth/verify", h.Verify)

    return mux
}
