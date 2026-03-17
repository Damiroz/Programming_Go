package main

import (
    "log"
    "net/http"
    "os"

    taskshttp "tech-ip-sem2/services/tasks/internal/http"
    "tech-ip-sem2/services/tasks/internal/grpcclient"
    "tech-ip-sem2/services/tasks/internal/service"
)

func main() {
    grpcAddr := os.Getenv("AUTH_GRPC_ADDR")
    if grpcAddr == "" {
        grpcAddr = "localhost:50051"
    }

    authClient, err := grpcclient.NewAuthClient(grpcAddr)
    if err != nil {
        log.Fatal("failed to connect to auth grpc:", err)
    }

    taskService := service.NewTaskService()
    handler := taskshttp.NewHandler(authClient, taskService)
    router := taskshttp.NewRouter(handler)

    port := os.Getenv("TASKS_PORT")
    if port == "" {
        port = "8082"
    }

    log.Println("Tasks HTTP server running on :" + port)
    http.ListenAndServe(":"+port, router)
}
