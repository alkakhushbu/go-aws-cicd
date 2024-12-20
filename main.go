package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/ping", ping)

	http.ListenAndServe(":80", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello ping! Pong has added some success test cases!"))
}
