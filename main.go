package main

import (
	"log"
	"net/http"

	"frontendmasters.com/reelingit/logger"
)

func main() {
	logInstance, err := logger.NewLogger("movies.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger %v", err)
	}
	defer logInstance.Close()

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed to initialize", err)
	}
}
