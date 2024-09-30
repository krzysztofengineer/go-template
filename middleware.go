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

			var color string
			switch {
			case rw.statusCode >= 100 && rw.statusCode <= 199:
				color = "\033[37m"
			case rw.statusCode >= 200 && rw.statusCode <= 299:
				color = "\033[32m"
			case rw.statusCode >= 300 && rw.statusCode <= 399:
				color = "\033[33m"
			case rw.statusCode >= 400 && rw.statusCode <= 499:
				color = "\033[34m"
			case rw.statusCode >= 500:
				color = "\033[31m"
			}

			slog.Info(fmt.Sprintf("%s%d\033[0m %s %s", color, rw.statusCode, r.Method, r.URL), "time=", time.Since(start))
		})
	}
}

func NewRecoverMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					slog.Error("panic", "error", r)

					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Internal Server Error"))
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
