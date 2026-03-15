package httpx

import (
    "context"
    "net"
    "net/http"
    "time"
)

func NewHTTPClient(timeout time.Duration) *http.Client {
    transport := &http.Transport{
        DialContext: (&net.Dialer{
            Timeout:   timeout,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        TLSHandshakeTimeout:   timeout,
        ResponseHeaderTimeout: timeout,
        ExpectContinueTimeout: 1 * time.Second,
    }

    return &http.Client{
        Transport: transport,
        Timeout:   timeout,
    }
}

func WithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
    return context.WithTimeout(ctx, timeout)
}
