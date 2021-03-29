package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var connLimiter = NewConnLimiter(2)

func main() {
	port := ":3000"
	var router = mux.NewRouter()
	router.Use(commonMiddleware)
	router.Use(rateLimitMiddleware)

	router.HandleFunc("/hello", handleMessage).Methods("GET")
	http.Handle("/", router)
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println("启动失败 ", err)
	}
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "{\"status\":\"ok\"}")
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !connLimiter.GetConn() {
			w.WriteHeader(500)
			fmt.Fprintf(w, "too many requests!")
			return
		}
		defer connLimiter.ReleaseConn()
		next.ServeHTTP(w, r)
	})
}
