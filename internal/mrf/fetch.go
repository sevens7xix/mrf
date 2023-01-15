package mrf

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sevens7xix/mrf/internal/model"
)

func getBearerToken() (string, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	params := url.Values{}
	params.Add("grant_type", `client_credentials`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)

	if err != nil {
		return "", fmt.Errorf("error crafting the token request '%s'", err)
	}

	encoded_credentials := base64.StdEncoding.EncodeToString([]byte("VIPER_IMPLEMENTATION:VIPER_IMPLEMENTATION"))

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded_credentials))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting a response from the Spotify API '%s'", err)
	}
	defer response.Body.Close()

	var result map[string]string

	json.NewDecoder(response.Body).Decode(&result)

	return result["access_token"], nil
}

func GetArtistID(artist []string) (string, error) {

	artist_query := stringFormatter(artist)

	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist", artist_query)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	bearerToken, err := getBearerToken()

	if err != nil {
		return "", err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var result model.SearchResponse

	json.NewDecoder(response.Body).Decode(&result)

	return result.Artists.Items[0].Href, nil
}
