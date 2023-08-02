package logging

import (
	"github.com/hotels-baby/go-adaptors/logging/factory"
	"github.com/hotels-baby/go-adaptors/logging/interfaces"
	"log"
)

type Config struct {
	LogFilePath     string
	LogLevel        string
	GoogleProjectID string // This is optional
}

type Client struct {
	logger interfaces.Logger
	config Config
}

func NewClient(t factory.LoggerType, c Config) *Client {

	logger, err := factory.NewLogger(t, c)
	if err != nil {
		log.Fatalf("Failed to load logger: %v", err)
	}

	return &Client{
		logger: logger,
		config: c,
	}
}
