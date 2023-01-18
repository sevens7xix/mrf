package mrf

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/sevens7xix/mrf/internal/utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetBearerToken(t *testing.T) {
	client := utilities.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`{"access_token":"Getting credentials..."}`)),
			Header:     make(http.Header),
		}
	})

	service := NewAlbumService([]string{"Phoenix"}, client)

	response, err := service.GetBearerToken()

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Getting credentials...", response)
}

func TestGetArtistIdHappyPath(t *testing.T) {
	client := utilities.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"artists": {
				  "href": "https://api.spotify.com/v1/search?query=michael+galaso&type=artist&locale=en-US%2Cen%3Bq%3D0.5&offset=0&limit=20",
				  "items": [
					{
					  "href": "https://api.spotify.com/v1/artists/0zwktRkdtjFn2F8v9kUdlF",
					  "id": "0zwktRkdtjFn2F8v9kUdlF",
					  "name": "Michael Galasso",
					  "uri": "spotify:artist:0zwktRkdtjFn2F8v9kUdlF"
					}
				  ],
				  "limit": 20,
				  "next": "https://api.spotify.com/v1/search?query=michael+galaso&type=artist&locale=en-US%2Cen%3Bq%3D0.5&offset=20&limit=20",
				  "offset": 0,
				  "previous": null,
				  "total": 36
				}
			  }`)),
			Header: make(http.Header),
		}
	})

	service := NewAlbumService([]string{"Michael", "Galasso"}, client)

	response, err := service.GetArtistID()

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "https://api.spotify.com/v1/artists/0zwktRkdtjFn2F8v9kUdlF", response)
}

func TestGetArtistIdUnhappyPath(t *testing.T) {
	client := utilities.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"artists": {
				  "href": "https://api.spotify.com/v1/search?query=michael+galaso&type=artist&locale=en-US%2Cen%3Bq%3D0.5&offset=0&limit=20",
				  "items": [],
				  "limit": 20,
				  "next": "https://api.spotify.com/v1/search?query=michael+galaso&type=artist&locale=en-US%2Cen%3Bq%3D0.5&offset=20&limit=20",
				  "offset": 0,
				  "previous": null,
				  "total": 36
				}
			  }`)),
			Header: make(http.Header),
		}
	})

	service := NewAlbumService([]string{"Mgwdfdfgjdn"}, client)

	response, err := service.GetArtistID()

	assert.Error(t, err)
	assert.Empty(t, response)
}
