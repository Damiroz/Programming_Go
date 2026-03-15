package main

import (
    "log"
    "net/http"
    "os"

    authhttp "tech-ip-sem2/services/auth/internal/http"
)

func main() {
    port := os.Getenv("AUTH_PORT")
    if port == "" {
        port = "8081"
    }

    router := authhttp.NewRouter()

    log.Printf("Auth service listening on :%s", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
        log.Fatalf("auth service failed: %v", err)
    }
}
