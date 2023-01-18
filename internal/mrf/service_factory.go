package mrf

import (
	"net/http"
)

func GetService(args []string, flag bool) (IService, error) {
	client := &http.Client{}

	if flag {
		return NewTopTrackService(args, client), nil
	} else {
		return NewAlbumService(args, client), nil
	}

}
