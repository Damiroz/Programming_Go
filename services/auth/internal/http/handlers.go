package http

import (
    "encoding/json"
    "net/http"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    json.NewDecoder(r.Body).Decode(&req)

    if req.Username == "student" && req.Password == "student" {
        json.NewEncoder(w).Encode(map[string]string{
            "access_token": "demo-token",
            "token_type":   "Bearer",
        })
        return
    }

    http.Error(w, "invalid credentials", http.StatusUnauthorized)
}
