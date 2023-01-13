package mrf

import (
	"io"
	"net/http"
)

type ServiceClient struct {
	Client *http.Client
}

func NewServiceClient(client *http.Client) *ServiceClient {
	return &ServiceClient{
		Client: client,
	}
}

func (s *ServiceClient) SendRequest(request *http.Request) (io.Reader, error) {
	response, err := s.Client.Do(request)

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}
