package mrf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequestHappyPath(t *testing.T) {
	service := NewAlbumService([]string{"Michael", "Galaso"}, client)

	request, err := service.createTracksRequest(
		"https://api.spotify.com/v1/artists/0zwktRkdtjFn2F8v9kUdlF/albums?market=US&limit=5")

	assert.NoError(t, err)
	assert.NotNil(t, request)
	assert.Equal(t,
		"https://api.spotify.com/v1/artists/0zwktRkdtjFn2F8v9kUdlF/albums?market=US&limit=5",
		request.URL.String())
}

func TestCreateRequestUnHappyPath(t *testing.T) {
	service := NewAlbumService([]string{"Michael", "Galaso"}, badClient)

	request, err := service.createTracksRequest("http//:wrong.request")

	assert.Error(t, err)
	assert.Nil(t, request)
}

func TestGetProcessedItems(t *testing.T) {
	service := NewAlbumService([]string{"Shakira"}, client)

	response, err := service.GetProcessedItems()

	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t,
		"=======================TRACKLIST=======================\n- Laundry Service: Washed and Dried (Expanded Edition).\n",
		response)
}
