package mrf

import (
	"net/http"
	"time"
)

func GetService(args []string, flag bool) (IService, error) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     10 * time.Second,
		},
		Timeout: 10 * time.Second,
	}

	if flag {
		return NewTopTrackService(args, client), nil
	} else {
		return NewAlbumService(args, client), nil
	}

}
