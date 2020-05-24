package godeezer

import (
	"time"
)

const (
	// DateLayout can be used with time.Parse to create time.Time values
	DateLayout = "2006-01-02"
)

// Album Deezer Album response
type Album struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Cover          string `json:"cover"`
	CoverSmall     string `json:"cover_small"`
	CoverMedium    string `json:"cover_medium"`
	CoverBig       string `json:"cover_big"`
	CoverXl        string `json:"cover_xl"`
	Tracklist      string `json:"tracklist"`
	Type           string `json:"type"`
	ReleaseDate    string `json:"release_date"`
	NUmberOfTracks int    `json:"nb_tracks"`
	Artist         Artist `json:"artist"`
}

// ReleaseDateTime get the release date time
func (s *Album) ReleaseDateTime() time.Time {
	result, _ := time.Parse(DateLayout, s.ReleaseDate)
	return result
}
