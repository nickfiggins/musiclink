package sptfy

type SpotifyImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type SpotifyExternalUrls struct {
	Spotify string `json:"spotify"`
} 


type SpotifyArtist struct {
	ExternalUrls SpotifyExternalUrls `json:"external_urls"`
	Href string `json:"href"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type SpotifyAlbum struct {
	AlbumType string `json:"album_type"`
	Artists   []SpotifyArtist `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	ExternalUrls SpotifyExternalUrls `json:"external_urls"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []SpotifyImage `json:"images"`
	Name                 string `json:"name"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks          int    `json:"total_tracks"`
	Type                 string `json:"type"`
	URI                  string `json:"uri"`
}