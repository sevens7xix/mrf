package model

type TopTracksResponse struct {
	Tracks []Tracks `json:"tracks"`
}

func (t TopTracksResponse) GetItems() []ProcessedItem {
	items := CastToProcessedItem(t.Tracks)
	return items
}

type Tracks struct {
	Album        TracksAlbum    `json:"album"`
	Artists      []TracksArtist `json:"artists"`
	DiscNumber   int            `json:"disc_number"`
	DurationMs   int            `json:"duration_ms"`
	Explicit     bool           `json:"explicit"`
	ExternalIds  External_Ids   `json:"external_ids"`
	ExternalUrls ExternalUrls   `json:"external_urls"`
	Href         string         `json:"href"`
	ID           string         `json:"id"`
	IsLocal      bool           `json:"is_local"`
	IsPlayable   bool           `json:"is_playable"`
	Name         string         `json:"name"`
	Popularity   int            `json:"popularity"`
	PreviewURL   string         `json:"preview_url"`
	TrackNumber  int            `json:"track_number"`
	Type         string         `json:"type"`
	URI          string         `json:"uri"`
}

func (t Tracks) GetName() string {
	return t.Name
}

type TracksAlbum struct {
	AlbumType            string         `json:"album_type"`
	Artists              []TracksArtist `json:"artists"`
	ExternalUrls         ExternalUrls   `json:"external_urls"`
	Href                 string         `json:"href"`
	ID                   string         `json:"id"`
	Images               []Image        `json:"images"`
	Name                 string         `json:"name"`
	ReleaseDate          string         `json:"release_date"`
	ReleaseDatePrecision string         `json:"release_date_precision"`
	TotalTracks          int            `json:"total_tracks"`
	Type                 string         `json:"type"`
	URI                  string         `json:"uri"`
}

type TracksArtist struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type External_Ids struct {
	Isrc string `json:"isrc"`
}
