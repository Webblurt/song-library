package services

import (
	"errors"
	models "song-library/internal/models"
	"time"
)

func (s *Service) FetchSongData(filters models.SongFilters) (models.SongLyrics, error) {
	lyrics, err := s.repository.GetSong(s.ctx, filters)
	if err != nil {
		s.log.Error("Error while fetching song data: ", err)
		return models.SongLyrics{}, err
	}
	return lyrics, nil
}

func (s *Service) SongCreation(req models.NewSongReq) error {
	songDetail, err := s.client.FetchSongInfo(req.Group, req.Song)
	if err != nil {
		s.log.Error("Error while fetching song info: ", err)
		return err
	}

	releaseDate, err := time.Parse("2006-01-02", songDetail.ReleaseDate)
	if err != nil {
		s.log.Error("Error while parsing release date: ", err)
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
		EntityName:       "songs",
		StringParameters: make(map[string]string),
		TimeParameters:   make(map[string]time.Time),
	}

	if req.Group != "" {
		entity.StringParameters["group"] = req.Group
	}
	if req.Song != "" {
		entity.StringParameters["song"] = req.Song
	}
	if req.Link != "" {
		entity.StringParameters["link"] = req.Link
	}
	if !req.ReleaseDate.IsZero() {
		entity.TimeParameters["release_date"] = req.ReleaseDate
	}

	if len(entity.StringParameters) == 0 && len(entity.TimeParameters) == 0 {
		s.log.Error("No fields to update")
		return errors.New("no fields to update")
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
