package services

import (
	models "song-library/internal/models"
	"time"
)

func (s *Service) FetchSongData(filters models.SongFilters) (models.SongLyrics, error) {
	lyrics, err := s.repository.GetSong(s.ctx, filters)
	if err != nil {
		return models.SongLyrics{}, err
	}
	return lyrics, nil
}

func (s *Service) SongCreation(req models.NewSongReq) error {
	songDetail, err := s.client.FetchSongInfo(req.Group, req.Song)
	if err != nil {
		return err
	}

	releaseDate, err := time.Parse("2006-01-02", songDetail.ReleaseDate)
	if err != nil {
		return err
	}

	entity := models.Entity{
		EntityName: "songs",
		StringParameters: map[string]string{
			"group": req.Group,
			"song":  req.Song,
			"link":  songDetail.Link,
		},
		TimeParameters: map[string]time.Time{
			"release_date": releaseDate,
		},
	}

	return s.repository.Create(s.ctx, entity)
}

func (s *Service) SongUpdating(req models.UpdateSongReq) error {
	entity := models.Entity{
		EntityName: "songs",
		StringParameters: map[string]string{
			"id":    req.ID,
			"group": req.Group,
			"song":  req.Song,
			"link":  req.Link,
		},
		TimeParameters: map[string]time.Time{
			"release_date": req.ReleaseDate,
		},
	}

	return s.repository.Update(s.ctx, entity)
}

func (s *Service) SongDeleting(req models.DeleteSongReq) error {
	entity := models.Entity{
		EntityName: "songs",
		StringParameters: map[string]string{
			"id": req.ID,
		},
	}

	return s.repository.Delete(s.ctx, entity)
}
