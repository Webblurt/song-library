package services

import (
	"context"
	"errors"
	clients "song-library/internal/clients"
	models "song-library/internal/models"
	repositories "song-library/internal/repositories"
	utils "song-library/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceInterface interface {
	FetchLibData(filters models.LibraryFilters) (models.Library, error)
	FetchSongData(filters models.SongFilters) (models.SongLyrics, error)
	SongCreation(req models.NewSongReq) error
	SongUpdating(req models.UpdateSongReq) error
	SongDeleting(req models.DeleteSongReq) error
}

type Service struct {
	db         *pgxpool.Pool
	repository *repositories.Repository
	client     *clients.ExternalAPIClient
	ctx        context.Context
	log        *utils.Logger
}

func NewService(client *clients.ExternalAPIClient, repo *repositories.Repository, log *utils.Logger) (*Service, error) {
	if repo == nil || client == nil {
		log.Error("Repository or client is nil")
		return nil, errors.New("repository or client is nil")
	}

	ctx := context.Background()

	return &Service{
		db:         repo.DB,
		repository: repo,
		client:     client,
		log:        log,
		ctx:        ctx,
	}, nil
}
