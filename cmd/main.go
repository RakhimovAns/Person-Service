package main

import (
	_ "github.com/RakhimovAns/Person-Service/docs"
	"github.com/RakhimovAns/Person-Service/internal/config"
	"github.com/RakhimovAns/Person-Service/internal/handler"
	"github.com/RakhimovAns/Person-Service/internal/repository"
	"github.com/RakhimovAns/Person-Service/internal/service"
	"github.com/RakhimovAns/Person-Service/pkg/client"
	"github.com/RakhimovAns/Person-Service/pkg/client/logging"
	"log"
)

// @title Person Service API
// @version 1.0
// @description API for managing people with data enrichment

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logging.New(cfg.LogLevel)

	db, err := repository.NewPostgresDB(cfg.DB)
	if err != nil {
		logger.Fatal("Failed to initialize db: %v", err)
	}

	personRepo := repository.NewPersonRepository(db, logger)
	agifyClient := client.NewAgifyClient(cfg.AgifyURL, logger)
	genderizeClient := client.NewGenderizeClient(cfg.GenderizeURL, logger)
	nationalizeClient := client.NewNationalizeClient(cfg.NationalizeURL, logger)

	personService := service.NewPersonService(personRepo, agifyClient, genderizeClient, nationalizeClient, logger)
	personHandler := handler.NewPersonHandler(personService, logger)

	server := handler.NewServer(cfg, personHandler)
	if err := server.Run(); err != nil {
		logger.Fatal("Server error: %v", err)
	}
}
