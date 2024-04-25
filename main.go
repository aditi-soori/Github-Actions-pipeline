package main

import (
	"fmt"

	"net/http"
)

func EndpointHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("message")
	if msg == "" {
		http.Error(w, "No message provided", http.StatusBadRequest)
		return
	}
	response := fmt.Sprintf("%s", msg)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/echo", EndpointHandler)

	http.ListenAndServe(":8080", nil)
}

// func Logger(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		rec := &responseLogger{w, http.StatusOK}
// 		handler.ServeHTTP(rec, r)
// 		log.Printf("%d", rec.status)
// 	}
// }

// type responseLogger struct {
// 	http.ResponseWriter
// 	status int
// }

// func (r *responseLogger) WriteHeader(statusCode int) {
// 	r.status = statusCode
// 	r.ResponseWriter.WriteHeader(statusCode)
// }
