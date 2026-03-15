package middleware

import (
    "context"
    "net/http"

    "github.com/google/uuid"
)

const (
    RequestIDHeader = "X-Request-ID"
    RequestIDKey    = "request_id"
)

func RequestID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        reqID := r.Header.Get(RequestIDHeader)
        if reqID == "" {
            reqID = uuid.New().String()
        }

        // Прокидываем в контекст
        ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
        r = r.WithContext(ctx)

        // Проставляем в ответ
        w.Header().Set(RequestIDHeader, reqID)

        next.ServeHTTP(w, r)
    })
}

func GetRequestID(r *http.Request) string {
    if v := r.Context().Value(RequestIDKey); v != nil {
        if s, ok := v.(string); ok {
            return s
        }
    }
    return ""
}
