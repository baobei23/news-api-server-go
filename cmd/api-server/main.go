package main

import (
	"log"
	"net/http"
	"news-api-server-go/router"
)

func main() {
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
