package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	traceID := r.Header.Get("X-B3-Traceid")
	spanID := r.Header.Get("X-B3-Spanid")
	log.Printf("X-B3-Traceid: %s", traceID)
	log.Printf("X-B3-Spanid: %s", spanID)
	req, err := http.NewRequest("GET", "http://backend.default", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to create the request: %v", err)
		return
	}

	req.Header.Set("X-B3-Traceid", traceID)
	req.Header.Set("X-B3-Spanid", spanID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to execute the request: %v", err)
		return
	}

	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func main() {
	log.Println("Starting...")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
