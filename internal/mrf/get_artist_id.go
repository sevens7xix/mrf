package mrf

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sevens7xix/mrf/internal/model"
	"github.com/sevens7xix/mrf/internal/utilities"
)

func (s *Service) GetBearerToken() (string, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	params := url.Values{}
	params.Add("grant_type", `client_credentials`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)

	if err != nil {
		return "", fmt.Errorf("error crafting the token request '%s'", err)
	}

	// We're going to get the credentials from the .env file, and later pass to the base64 encoder
	credentials, err := utilities.GetCredentials()

	if err != nil {
		return "", fmt.Errorf("problem getting the app credentials '%s'", err)
	}

	credentials_string := fmt.Sprintf("%s:%s", credentials[0], credentials[1])

	encoded_credentials := base64.StdEncoding.EncodeToString([]byte(credentials_string))

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded_credentials))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := s.Client.Do(req)

	if err != nil {
		return "", fmt.Errorf("error sending the token request '%s'", err)
	}

	defer response.Body.Close()

	var result map[string]string

	json.NewDecoder(response.Body).Decode(&result)

	return result["access_token"], nil
}

func (s *Service) GetArtistID(artist []string) (string, error) {

	artist_query := utilities.StringFormatter(artist)

	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist", artist_query)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	bearerToken, err := s.GetBearerToken()

	if err != nil {
		return "", err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	response, err := s.Client.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var result model.SearchResponse

	json.NewDecoder(response.Body).Decode(&result)

	if len(result.Artists.Items) == 0 {
		return "", fmt.Errorf("'%d' artist found for this query", len(result.Artists.Items))
	}

	return result.Artists.Items[0].Href, nil
}