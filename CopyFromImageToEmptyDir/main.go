package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var path = "/var/log/app.log"

func main() {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open %q: %v'", path, err)
	}

	for {
		_, err := fmt.Fprint(f, "This is a log from the app!\n")
		if err != nil {
			log.Fatalf("failed to write to %q: %v", path, err)
		}
		time.Sleep(1 * time.Second)
	}
}
