package godeezer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Artist Deezer Artist response
type Artist struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Type          string `json:"type"`
}

type artistTracklist struct {
	Data []Track `json:"data"`
}

var (
	deezerSearchAPIURL = "https://api.deezer.com/search/artist?q=%s"
	deezerSearchWebURL = "https://www.deezer.com/search/%s/artist"
)

// GetTopTracks for the currrent Artist
func (a Artist) GetTopTracks(limit int) ([]Track, error) {
	return GetTopTracksForArtistID(a.ID, limit)
}

// SearchForArtistViaAPI return the matched Artist using the Deezer API
func SearchForArtistViaAPI(artistName string) (*Artist, error) {

	type artistSearchResult struct {
		Data    []Artist `json:"data"`
		NextURL string   `json:"next,omitempty"`
	}

	artistName = strings.ToLower(artistName)
	var result *Artist
	url := fmt.Sprintf(deezerSearchAPIURL, url.QueryEscape(artistName))
	res, httpErr := http.Get(url)
	if httpErr != nil {
		return result, httpErr
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return result, fmt.Errorf("Status code was %d", res.StatusCode)
	}

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return result, bodyErr
	}

	searchResult := artistSearchResult{}
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return result, err
	}

	if len(searchResult.Data) == 1 {
		return &searchResult.Data[0], nil
	}
	for _, artist := range searchResult.Data {
		if strings.EqualFold(artist.Name, artistName) {
			return &artist, nil
		}
	}

	return result, nil
}

//SearchForArtistIDViaWeb Use the Deezer Web Search service
func SearchForArtistIDViaWeb(artistName string) (int, error) {
	type deezerWebSearch struct {
		TOPRESULT []struct {
			ARTID   string `json:"ART_ID"`
			ARTNAME string `json:"ART_NAME"`
		} `json:"TOP_RESULT"`
	}

	artistName = strings.ToLower(artistName)
	var result int
	url := fmt.Sprintf(deezerSearchWebURL, artistName)
	res, httpErr := http.Get(url)
	if httpErr != nil {
		return result, httpErr
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return result, fmt.Errorf("Status code was %d", res.StatusCode)
	}

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return result, bodyErr
	}

	re := regexp.MustCompile(`<script>window.__DZR_APP_STATE__\s* =\s*(.*)<\/script>`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) == 0 {
		return result, nil
	}

	searchResult := deezerWebSearch{}
	if err := json.Unmarshal([]byte(matches[1]), &searchResult); err != nil {
		return result, err
	}

	if len(searchResult.TOPRESULT) > 0 {
		return strconv.Atoi(searchResult.TOPRESULT[0].ARTID)
	}

	return result, nil
}
