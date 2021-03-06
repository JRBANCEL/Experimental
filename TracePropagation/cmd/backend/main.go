package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Headers: %#v", r.Header)
	fmt.Fprintf(w, "From Backend with love.\n")
}

func main() {
	log.Println("Starting...")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
