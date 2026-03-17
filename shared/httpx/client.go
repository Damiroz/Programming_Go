package httpx

import (
    "net/http"
    "time"
)

func NewClient() *http.Client {
    return &http.Client{
        Timeout: 3 * time.Second,
    }
}
