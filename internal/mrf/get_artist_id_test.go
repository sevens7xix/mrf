package mrf

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/sevens7xix/mrf/internal/utilities"
	"github.com/stretchr/testify/assert"
)

var client, badClient *http.Client

func TestMain(m *testing.M) {

	client = utilities.NewTestClient(func(req *http.Request) *http.Response {

		request_string := req.URL.String()

		if request_string == "https://accounts.spotify.com/api/token" {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString(`{"access_token":"Getting credentials..."}`)),
				Header:     make(http.Header),
			}
		} else if strings.Contains(request_string, "albums?market=US&limit=5") {
			return &http.Response{
				StatusCode: 404,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"href": "https://api.spotify.com/v1/artists/0EmeFodog0BfCgMzAIvKQp/albums?offset=0&limit=20&include_groups=album,single,compilation,appears_on&locale=en-US,en;q=0.5",
					"items": [
					  {
						"album_group": "album",
						"album_type": "album",
						"artists": [
						  {
							"external_urls": {
							  "spotify": "https://open.spotify.com/artist/0EmeFodog0BfCgMzAIvKQp"
							},
							"href": "https://api.spotify.com/v1/artists/0EmeFodog0BfCgMzAIvKQp",
							"id": "0EmeFodog0BfCgMzAIvKQp",
							"name": "Shakira",
							"type": "artist",
							"uri": "spotify:artist:0EmeFodog0BfCgMzAIvKQp"
						  }
						],
						"external_urls": {
						  "spotify": "https://open.spotify.com/album/6WaruQqgJzSlSzZz2YdUku"
						},
						"href": "https://api.spotify.com/v1/albums/6WaruQqgJzSlSzZz2YdUku",
						"id": "6WaruQqgJzSlSzZz2YdUku",
						"name": "Laundry Service: Washed and Dried (Expanded Edition)",
						"release_date": "2021-11-12",
						"release_date_precision": "day",
						"total_tracks": 17,
						"type": "album",
						"uri": "spotify:album:6WaruQqgJzSlSzZz2YdUku"
					  }
					],
					"limit": 20,
					"next": "https://api.spotify.com/v1/artists/0EmeFodog0BfCgMzAIvKQp/albums?offset=20&limit=20&include_groups=album,single,compilation,appears_on&locale=en-US,en;q=0.5",
					"offset": 0,
					"previous": null,
					"total": 548
				  }`)),
				Header: make(http.Header),
			}
		} else {
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
		}
	})

	badClient = utilities.NewTestClient(func(req *http.Request) *http.Response {
		if req.URL.String() == "https://accounts.spotify.com/api/token" {
			return &http.Response{
				StatusCode: 404,
				Body:       io.NopCloser(bytes.NewBufferString(`{"error":"invalid_client"}`)),
				Header:     make(http.Header),
			}
		} else {
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
		}
	})

	code := m.Run()
	os.Exit(code)
}

func TestGetBearerTokenHappyPath(t *testing.T) {
	service := NewAlbumService([]string{"Phoenix"}, client)

	response, err := service.GetBearerToken()

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Getting credentials...", response)
}

func TestGetBearerTokenUnHappyPath(t *testing.T) {
	service := NewAlbumService([]string{"Phoenix"}, badClient)

	response, err := service.GetBearerToken()

	assert.Error(t, err)
	assert.Empty(t, response)
}

func TestGetArtistIdHappyPath(t *testing.T) {
	service := NewAlbumService([]string{"Michael", "Galasso"}, client)

	response, err := service.GetArtistID()

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "https://api.spotify.com/v1/artists/0zwktRkdtjFn2F8v9kUdlF", response)
}

func TestGetArtistIdUnhappyPath(t *testing.T) {
	service := NewAlbumService([]string{"Mgwdfdfgjdn"}, badClient)

	response, err := service.GetArtistID()

	assert.Error(t, err)
	assert.Empty(t, response)
}
