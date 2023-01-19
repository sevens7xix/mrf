package mrf

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sevens7xix/mrf/internal/utilities"
)

func (s *Service) createTracksRequest(URL string) (*http.Request, error) {

	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		return nil, err
	}

	bearerToken, err := s.GetBearerToken()

	if err != nil {
		return nil, fmt.Errorf("an error happened with the token '%s'", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	return req, nil
}

func (s *Service) GetProcessedItems() (string, error) {
	artistID, err := s.GetArtistID()

	if err != nil {
		return "", fmt.Errorf("something wrong happened geting the artist ID '%s'", err)
	}

	URL := fmt.Sprintf("%s/%s", artistID, s.URI)

	req, err := s.createTracksRequest(URL)

	if err != nil {
		return "", fmt.Errorf("something wrong happened creating the API request '%s'", err)
	}

	response, err := s.Client.Do(req)

	if err != nil {
		return "", fmt.Errorf("something wrong happened getting the API response '%s'", err)
	}

	err = json.NewDecoder(response.Body).Decode(&s.Model)

	if err != nil {
		return "", fmt.Errorf("something wrong happened decoding the response '%s'", err)
	}

	track_names := utilities.ProcessTitleTracks(s.Model.GetItems())

	return track_names, nil
}
