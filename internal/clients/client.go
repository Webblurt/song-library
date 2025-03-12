package clients

import (
	models "song-library/internal/models"
	utils "song-library/internal/utils"
)

type APIClient interface {
	FetchSongInfo(baseURL, group, song string) (models.SongDetail, error)
}

type ExternalAPIClient struct {
	Name string
	URL  string
	log  *utils.Logger
}

func NewExternalAPIClient(cfg *utils.Config, log *utils.Logger) (*ExternalAPIClient, error) {
	return &ExternalAPIClient{
		Name: cfg.API.Name,
		URL:  cfg.API.URL,
		log:  log,
	}, nil
}
