package main

import (
	_ "github.com/e-wrobel/swagger/docs"
	"github.com/e-wrobel/swagger/internal/handlers"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "experiments", log.LstdFlags)
	service := handlers.NewService(logger)
	logger.Printf("Service initialized: %v", service)
}
