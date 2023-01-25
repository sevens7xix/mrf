package mrf

import (
	"net/http"

	"github.com/sevens7xix/mrf/internal/model"
)

// IService is the interface that works and abstraction for the service struct and the service factory
type IService interface {
	GetBearerToken() (string, error)
	GetArtistID() (string, error)
	createTracksRequest(URL string) (*http.Request, error)
	GetProcessedItems() (string, error)
}

// Service is the base implementation of the IService interface and the base of the factory implementation
type Service struct {
	Artist []string
	URI    string
	Client *http.Client
	Model  model.ProcessedResponse
}
