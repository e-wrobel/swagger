package main

import (
	"github.com/e-wrobel/swagger/internal/handlers"
	_ "github.com/pdrum/swagger-automation/docs"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "experiments", log.LstdFlags)
	service := handlers.NewService(logger)
	logger.Printf("Service initialized: %v", service)
}
