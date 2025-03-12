package models

type Song struct {
	ID          string `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
}

type Library struct {
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	LastPage int    `json:"last_page"`
	HasNext  bool   `json:"has_next"`
	HasPrev  bool   `json:"has_prev"`
	NextPage int    `json:"next_page"`
	PrevPage int    `json:"prev_page"`
	Songs    []Song `json:"songs"`
}

type SongLyrics struct {
	ID          string `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	Title       string `json:"title"`
	Lyrics      string `json:"lyrics"`
	Page        int    `json:"page"`
	TotalVerses int    `json:"total_verses"`
	HasNext     bool   `json:"has_next"`
	HasPrev     bool   `json:"has_prev"`
	NextPage    int    `json:"next_page,omitempty"`
	PrevPage    int    `json:"prev_page,omitempty"`
}
