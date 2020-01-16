package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Failed to proxy the request: %v", err)
	} else {
		io.Copy(w, resp.Body)
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
