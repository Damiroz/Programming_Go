package main

import (
    "log"
    "net/http"
    "os"

    authgrpc "tech-ip-sem2/services/auth/internal/grpc"
    authhttp "tech-ip-sem2/services/auth/internal/http"
)

func main() {
    grpcPort := os.Getenv("AUTH_GRPC_PORT")
    if grpcPort == "" {
        grpcPort = "50051"
    }

    httpPort := os.Getenv("AUTH_HTTP_PORT")
    if httpPort == "" {
        httpPort = "8081"
    }

    // gRPC server
    go func() {
        if err := authgrpc.RunGRPC(":" + grpcPort); err != nil {
            log.Fatal(err)
        }
    }()

    // HTTP login server
    log.Println("Auth HTTP server running on :" + httpPort)
    http.ListenAndServe(":"+httpPort, authhttp.NewRouter())
}
