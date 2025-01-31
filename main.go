package main

import (
	"experiments/internal/handlers"
	_ "g/docs"
	_ "github.com/pdrum/swagger-automation/docs"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "experiments", log.LstdFlags)
	service := handlers.NewService(logger)
	logger.Printf("Service initialized: %v", service)
}
