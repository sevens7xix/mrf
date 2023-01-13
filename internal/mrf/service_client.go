package mrf

import (
	"io"
	"net/http"
)

type Service_Client interface {
	SendRequest(request *http.Request) (io.Reader, error)
}
