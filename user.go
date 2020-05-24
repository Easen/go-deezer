package godeezer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	deezerUserArtists = "https://api.deezer.com/user/me/artists?access_token=%s"
	deezerUserAlbums  = "https://api.deezer.com/user/me/albums?access_token=%s"
	deezerUserTracks  = "https://api.deezer.com/user/me/tracks?access_token=%s"
)

// GetUserFavoriteArtists return the users favorite artists
func GetUserFavoriteArtists(deezerAccessToken string) ([]Artist, error) {
	var artists []Artist
	var url = fmt.Sprintf(deezerUserArtists, deezerAccessToken)
	for {
		var result struct {
			Data    []Artist `json:"data"`
			Total   int      `json:"total"`
			NextURL *string  `json:"next,omitempty"`
		}
		err := getUrl(url, &result)
		if err != nil {
			return nil, err
		}
		artists = append(artists, result.Data...)
		if result.NextURL == nil {
			break
		}
		url = *result.NextURL
	}
	return artists, nil
}

// GetUserFavoriteAlbums return the users favorite albums
func GetUserFavoriteAlbums(deezerAccessToken string) ([]Album, error) {
	var albums []Album
	var url = fmt.Sprintf(deezerUserAlbums, deezerAccessToken)
	for {
		var result struct {
			Data    []Album `json:"data"`
			Total   int     `json:"total"`
			NextURL *string `json:"next,omitempty"`
		}
		err := getUrl(url, &result)
		if err != nil {
			return nil, err
		}
		albums = append(albums, result.Data...)
		if result.NextURL == nil {
			break
		}
		url = *result.NextURL
	}
	return albums, nil
}

// GetUserFavoriteTracks return the users favorite albums
func GetUserFavoriteTracks(deezerAccessToken string) ([]Track, error) {
	var tracks []Track
	var url = fmt.Sprintf(deezerUserTracks, deezerAccessToken)
	for {
		var result struct {
			Data    []Track `json:"data"`
			Total   int     `json:"total"`
			NextURL *string `json:"next,omitempty"`
		}
		err := getUrl(url, &result)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, result.Data...)
		if result.NextURL == nil {
			break
		}
		url = *result.NextURL
	}
	return tracks, nil
}

func getUrl(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Non success status code %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
