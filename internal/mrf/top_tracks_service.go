package mrf

import (
	"net/http"

	"github.com/sevens7xix/mrf/internal/model"
)

type TopTrackService struct {
	Service
}

func NewTopTrackService(artist []string, client *http.Client) *TopTrackService {
	return &TopTrackService{
		Service: Service{
			Artist: artist,
			URI:    "top-tracks?market=US",
			Client: client,
			Model:  &model.TopTracksResponse{},
		},
	}
}
