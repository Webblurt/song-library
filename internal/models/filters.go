package models

type SongFilters struct {
	SongID string `json:"song_id"`
	Group  string `json:"group"`
	Song   string `json:"song"`
	Page   string `json:"page"`
	Limit  string `json:"limit"`
}

type LibraryFilters struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Limit       string `json:"limit"`
	Offset      string `json:"offset"`
}
