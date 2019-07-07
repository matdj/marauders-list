package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/crossoff", makeHandler(crossoffHandler))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
