package logging

import (
	"github.com/hotels-baby/go-adaptors/logging/factory"
	"github.com/hotels-baby/go-adaptors/logging/interfaces"
	"log"
)

type Client struct {
	logger interfaces.Logger
}

func NewClient(t factory.LoggerType, c factory.LoggerConfig) *Client {

	logger, err := factory.NewLogger(t, c)
	if err != nil {
		log.Fatalf("Failed to load logger: %v", err)
	}

	return &Client{
		logger: logger,
	}
}
