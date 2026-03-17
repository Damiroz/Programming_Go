package http

import "net/http"

func NewRouter() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/v1/auth/login", LoginHandler)
    return mux
}
