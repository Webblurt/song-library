package services

import (
	models "song-library/internal/models"
)

func (s *Service) FetchLibData(filters models.LibraryFilters) (models.Library, error) {
	library, err := s.repository.GetLibrary(s.ctx, filters)
	if err != nil {
		return models.Library{}, err
	}
	return library, nil
}
