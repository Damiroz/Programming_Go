package authclient

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"

    "tech-ip-sem2/shared/httpx"
    "tech-ip-sem2/shared/middleware"
)

type VerifyResponse struct {
    Valid   bool   `json:"valid"`
    Subject string `json:"subject,omitempty"`
    Error   string `json:"error,omitempty"`
}

type Client struct {
    baseURL string
    client  *http.Client
}

func NewClient() *Client {
    baseURL := os.Getenv("AUTH_BASE_URL")
    if baseURL == "" {
        baseURL = "http://localhost:8081"
    }

    return &Client{
        baseURL: baseURL,
        client:  httpx.NewHTTPClient(3 * time.Second),
    }
}

func (c *Client) Verify(ctx context.Context, token string, reqID string) (VerifyResponse, int, error) {
    ctx, cancel := httpx.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/v1/auth/verify", nil)
    if err != nil {
        return VerifyResponse{}, 0, fmt.Errorf("create request: %w", err)
    }

    req.Header.Set("Authorization", "Bearer "+token)
    if reqID != "" {
        req.Header.Set(middleware.RequestIDHeader, reqID)
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return VerifyResponse{}, 0, fmt.Errorf("do request: %w", err)
    }
    defer resp.Body.Close()

    var vr VerifyResponse
    if err := json.NewDecoder(resp.Body).Decode(&vr); err != nil {
        return VerifyResponse{}, resp.StatusCode, fmt.Errorf("decode response: %w", err)
    }

    return vr, resp.StatusCode, nil
}
