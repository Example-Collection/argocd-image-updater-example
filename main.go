package main

import "net/http"

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message": "server is running, version 0.0.4"}`))
	})
	_ = http.ListenAndServe(":8080", nil)
}
