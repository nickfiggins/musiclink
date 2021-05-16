package sptfy

import (
	"encoding/json"
	"github.com/rapito/go-spotify/spotify"
	"os"
)


type SpotifyAPI struct {
	Spot spotify.Spotify
}


func NewSpotifyAPI() *SpotifyAPI{
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	spot := spotify.New(clientID,clientSecret)
	spot.Authorize()
	return &SpotifyAPI{Spot: spot}
}


func (s SpotifyAPI) GetTrack(id string) (SpotifyTrackEndpoint, error) {
	s.Spot.Authorize()
	var track SpotifyTrackEndpoint
	bytes, _ := s.Spot.Get("tracks/%s", nil, id)
	json.Unmarshal(bytes, &track)
	return track, nil
}

func (s SpotifyAPI) GetAlbum(id string) (SpotifyAlbumEndpoint, error) {
	s.Spot.Authorize()
	var album SpotifyAlbumEndpoint
	bytes, _ := s.Spot.Get("albums/%s", nil, id)
	json.Unmarshal(bytes, &album)
	return album, nil
}

func (s SpotifyAPI) GetArtist(id string) (SpotifyArtistEndpoint, error) {
	s.Spot.Authorize()
	var artist SpotifyArtistEndpoint
	bytes, _ := s.Spot.Get("artists/%s", nil, id)
	json.Unmarshal(bytes, &artist)
	return artist, nil
}