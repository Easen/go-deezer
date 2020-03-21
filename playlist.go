package godeezer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	deezerPlaylistURL = "https://api.deezer.com/playlist/%s/tracks?request_method=POST&access_token=%s&songs=%s"
)

//UpdatePlaylistTracks Update the Tracks of a Deezer playlist
func UpdatePlaylistTracks(deezerAccessToken string, playlistID int, trackIDs []int) error {
	var songs strings.Builder
	for _, trackID := range trackIDs {
		if songs.Len() > 0 {
			fmt.Fprintf(&songs, ",")
		}
		fmt.Fprintf(&songs, "%d", trackID)

	}

	updatePlaylistTracksURL := fmt.Sprintf(deezerPlaylistURL, strconv.Itoa(playlistID), deezerAccessToken, songs.String())

	fmt.Print(updatePlaylistTracksURL)
	res, err := http.Get(updatePlaylistTracksURL)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("Returned status code as %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Printf("Received body: %s\n", string(body))

	return nil
}
