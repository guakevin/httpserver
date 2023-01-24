package main

import (
	"cmd/api/main.go/internal/app"
	"log"
)

func main() {
	c := app.APIConfig{Host: "localhost", Port: 8000}
	server := app.New(&c)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
