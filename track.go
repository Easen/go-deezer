package godeezer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Track - A Deezer Track resource
type Track struct {
	ID                    int    `json:"id"`
	Readable              bool   `json:"readable"`
	Title                 string `json:"title"`
	TitleShort            string `json:"title_short"`
	TitleVersion          string `json:"title_version"`
	Link                  string `json:"link"`
	Duration              int    `json:"duration"`
	Rank                  int    `json:"rank"`
	ExplicitLyrics        bool   `json:"explicit_lyrics"`
	ExplicitContentLyrics int    `json:"explicit_content_lyrics"`
	ExplicitContentCover  int    `json:"explicit_content_cover"`
	Preview               string `json:"preview"`
	Artist                struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Link          string `json:"link"`
		Picture       string `json:"picture"`
		PictureSmall  string `json:"picture_small"`
		PictureMedium string `json:"picture_medium"`
		PictureBig    string `json:"picture_big"`
		PictureXl     string `json:"picture_xl"`
		Tracklist     string `json:"tracklist"`
		Type          string `json:"type"`
	} `json:"artist"`
	Album struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Cover       string `json:"cover"`
		CoverSmall  string `json:"cover_small"`
		CoverMedium string `json:"cover_medium"`
		CoverBig    string `json:"cover_big"`
		CoverXl     string `json:"cover_xl"`
		Tracklist   string `json:"tracklist"`
		Type        string `json:"type"`
	} `json:"album"`
	Type string `json:"type"`
}

var (
	deezerTopTracks = "https://api.deezer.com/artist/%d/top?limit=%d"
)

// GetTopTracksForArtistID Get the top tracks for an ArtistID
func GetTopTracksForArtistID(artistID int, limit int) ([]Track, error) {
	var tracks []Track
	requestURL := fmt.Sprintf(deezerTopTracks, artistID, limit)
	res, httpErr := http.Get(requestURL)
	if httpErr != nil {
		return tracks, httpErr
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return tracks, fmt.Errorf("Status code was %d", res.StatusCode)
	}

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return tracks, bodyErr
	}

	var result artistTracklist
	if err := json.Unmarshal(body, &result); err != nil {
		return tracks, err
	}

	if len(result.Data) <= limit {
		return result.Data, nil
	}

	return result.Data[0:limit], nil
}
