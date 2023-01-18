package mrf

import (
	"fmt"
	"net/http"
)

func (s *Service) createTracksRequest(URL string) (*http.Request, error) {

	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		return nil, err
	}

	bearerToken, err := s.GetBearerToken()

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	return req, nil
}
