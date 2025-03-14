package repositories

import (
	"context"
	"fmt"
	models "song-library/internal/models"
	"strconv"
)

func (r *Repository) GetLibrary(ctx context.Context, filters models.LibraryFilters) (models.Library, error) {
	query := `
		SELECT s.id, g.name AS group_name, s.song, s.release_date 
		FROM songs s
		JOIN groups g ON s.group_id = g.id
		WHERE 1=1`
	countQuery := `
		SELECT COUNT(*) 
		FROM songs s
		JOIN groups g ON s.group_id = g.id
		WHERE 1=1`
	args := []interface{}{}
	countArgs := []interface{}{}
	argID := 1
	countArgID := 1

	if filters.Group != "" {
		query += fmt.Sprintf(" AND g.name ILIKE $%d", argID)
		countQuery += fmt.Sprintf(" AND g.name ILIKE $%d", countArgID)
		args = append(args, "%"+filters.Group+"%")
		countArgs = append(countArgs, "%"+filters.Group+"%")
		argID++
		countArgID++
	}
	if filters.Song != "" {
		query += fmt.Sprintf(" AND s.song ILIKE $%d", argID)
		countQuery += fmt.Sprintf(" AND s.song ILIKE $%d", countArgID)
		args = append(args, "%"+filters.Song+"%")
		countArgs = append(countArgs, "%"+filters.Song+"%")
		argID++
		countArgID++
	}
	if filters.ReleaseDate != "" {
		query += fmt.Sprintf(" AND s.release_date = $%d", argID)
		countQuery += fmt.Sprintf(" AND s.release_date = $%d", countArgID)
		args = append(args, filters.ReleaseDate)
		countArgs = append(countArgs, filters.ReleaseDate)
		argID++
		countArgID++
	}

	var total int
	err := r.DB.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return models.Library{}, err
	}

	limit, err := strconv.Atoi(filters.Limit)
	if err != nil {
		return models.Library{}, err
	}
	offset, err := strconv.Atoi(filters.Offset)
	if err != nil {
		return models.Library{}, err
	}

	query += fmt.Sprintf(" ORDER BY s.release_date DESC LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return models.Library{}, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var s models.Song
		err := rows.Scan(&s.ID, &s.Group, &s.Song, &s.ReleaseDate)
		if err != nil {
			return models.Library{}, err
		}
		songs = append(songs, s)
	}

	page := offset/limit + 1
	lastPage := (total + limit - 1) / limit
	hasNext := page < lastPage
	hasPrev := page > 1
	nextPage := page + 1
	if !hasNext {
		nextPage = 0
	}
	prevPage := page - 1
	if !hasPrev {
		prevPage = 0
	}

	return models.Library{
		Total:    total,
		Page:     page,
		PerPage:  limit,
		LastPage: lastPage,
		HasNext:  hasNext,
		HasPrev:  hasPrev,
		NextPage: nextPage,
		PrevPage: prevPage,
		Songs:    songs,
	}, nil
}
