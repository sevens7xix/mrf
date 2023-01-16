package mrf

import (
	"net/http"
)

type Service struct {
	Client *http.Client
}

func NewService(client *http.Client) *Service {
	return &Service{
		Client: client,
	}
}
