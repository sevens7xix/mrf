package mrf

import (
	"fmt"
	"net/http"
)

func (s *Service) getAlbumsRequest(artistURL string) (*http.Request, error) {
	newURL := fmt.Sprintf("%s/albums?market=US&limit=5", artistURL)

	req, err := http.NewRequest("GET", newURL, nil)

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
