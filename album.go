package godeezer

// Album Deezer Album response
type Album struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	CoverSmall  string `json:"cover_small"`
	CoverMedium string `json:"cover_medium"`
	CoverBig    string `json:"cover_big"`
	CoverXl     string `json:"cover_xl"`
	Tracklist   string `json:"tracklist"`
	Type        string `json:"type"`
}
