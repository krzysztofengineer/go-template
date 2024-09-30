package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func NewMiddlewareStack(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next
	}
}

func NewLoggerMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := NewResponseWriter(w)

			next.ServeHTTP(rw, r)

			slog.Info(fmt.Sprintf("%d %s %s", rw.statusCode, r.Method, r.URL), "time=", time.Since(start))
		})
	}
}
