package main

import (
	"cmd/api/main.go/internal/app"

	"flag"
	"log"
)

func main() {
	port := flag.Int("p", 8000, "PORT")
	logLevel := flag.String("l", "debug", "LOG LEVEL")
	flag.Parse()

	c := app.APIConfig{Host: "0.0.0.0", Port: *port, LogLevel: *logLevel}
	server := app.New(c)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
