package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/", index)
	log.Println("Starting service on port 8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("/")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("/health")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy\n"))
}
