package mrf

import (
	"net/http"

	"github.com/sevens7xix/mrf/internal/model"
)

type IService interface {
	GetBearerToken() (string, error)
	GetArtistID() (string, error)
	createTracksRequest(URL string) (*http.Request, error)
	GetProcessedItems() (string, error)
}

type Service struct {
	Artist []string
	URI    string
	Client *http.Client
	Model  model.ProcessedResponse
}
