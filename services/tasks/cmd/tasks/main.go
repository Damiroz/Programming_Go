package main

import (
    "log"
    "net/http"
    "os"

    taskhttp "tech-ip-sem2/services/tasks/internal/http"
)

func main() {
    port := os.Getenv("TASKS_PORT")
    if port == "" {
        port = "8082"
    }

    router := taskhttp.NewRouter()

    log.Printf("Tasks service listening on :%s", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
        log.Fatalf("tasks service failed: %v", err)
    }
}
