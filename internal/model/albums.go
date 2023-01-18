package model

type Albums struct {
	Href     string      `json:"href"`
	Items    []AlbumItem `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous string      `json:"previous"`
	Total    int         `json:"total"`
}

type AlbumItem struct {
	AlbumGroup           string           `json:"album_group"`
	AlbumType            string           `json:"album_type"`
	Artists              []AlbumArtists   `json:"artists"`
	AvailableMarkets     []string         `json:"available_markets"`
	ExternalUrls         AlbumExternalUrl `json:"external_urls"`
	Href                 string           `json:"href"`
	ID                   string           `json:"id"`
	Images               []AlbumImages    `json:"images"`
	Name                 string           `json:"name"`
	ReleaseDate          string           `json:"release_date"`
	ReleaseDatePrecision string           `json:"release_date_precision"`
	TotalTracks          int              `json:"total_tracks"`
	Type                 string           `json:"type"`
	URI                  string           `json:"uri"`
}

func (a AlbumItem) GetName() string {
	return a.Name
}

type AlbumArtists struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type AlbumImages struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type AlbumExternalUrl struct {
	Spotify string `json:"spotify"`
}
