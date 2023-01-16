package mrf

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBearerToken(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"access_token":"Getting credentials..."}`)),
			Header:     make(http.Header),
		}
	})

	service := NewService(client)

	response, err := service.GetBearerToken()

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Getting credentials...", response)
}
