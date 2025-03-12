package models

import "time"

type NewSongReq struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type UpdateSongReq struct {
	ID          string    `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate time.Time `json:"release_date"`
	Link        string    `json:"link"`
}

type DeleteSongReq struct {
	ID string `json:"id"`
}
