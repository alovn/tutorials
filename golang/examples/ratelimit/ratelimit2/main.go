package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(2, 5)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			var httpCode = http.StatusTooManyRequests
			http.Error(w, http.StatusText(httpCode), httpCode)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	http.ListenAndServe(":3000", limit(mux))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
