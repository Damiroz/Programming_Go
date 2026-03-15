package middleware

import (
    "log"
    "net/http"
    "time"
)

type loggingResponseWriter struct {
    http.ResponseWriter
    status int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.status = code
    lrw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        reqID := GetRequestID(r)

        lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}

        next.ServeHTTP(lrw, r)

        log.Printf("[request_id=%s] %s %s %d %s",
            reqID,
            r.Method,
            r.URL.Path,
            lrw.status,
            time.Since(start),
        )
    })
}
