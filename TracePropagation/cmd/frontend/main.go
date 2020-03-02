package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s", r.RemoteAddr)
	resp, err := http.Get("http://backend.default")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	} else {
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	}
}

func main() {
	log.Println("Starting...")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
