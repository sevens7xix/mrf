package mrf

import (
	"net/http"

	"github.com/sevens7xix/mrf/internal/model"
)

type AlbumService struct {
	Service
}

func NewAlbumService(artist []string, client *http.Client) *AlbumService {
	return &AlbumService{
		Service: Service{
			Artist: artist,
			URI:    "albums?market=US&limit=5",
			Client: client,
			Model:  &model.Albums{},
		},
	}
}
