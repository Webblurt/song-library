package repositories

import (
	"context"
	"fmt"
	models "song-library/internal/models"
	"strconv"
)

func (r *Repository) GetSong(ctx context.Context, filters models.SongFilters) (models.SongLyrics, error) {
	query := `
		SELECT l.text, g.name AS group_name, s.song
		FROM lyrics l
		JOIN songs s ON l.song_id = s.id
		JOIN groups g ON s.group_id = g.id
		WHERE 1=1`
	args := []interface{}{}
	argID := 1

	if filters.SongID != "" {
		query += fmt.Sprintf(" AND s.id = $%d", argID)
		args = append(args, filters.SongID)
		argID++
	} else {
		if filters.Group != "" {
			query += fmt.Sprintf(" AND g.name ILIKE $%d", argID)
			args = append(args, "%"+filters.Group+"%")
			argID++
		}
		if filters.Song != "" {
			query += fmt.Sprintf(" AND s.song ILIKE $%d", argID)
			args = append(args, "%"+filters.Song+"%")
			argID++
		}
	}

	query += " ORDER BY l.verse_number"

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return models.SongLyrics{}, err
	}
	defer rows.Close()

	var verses []string
	for rows.Next() {
		var verse string
		err := rows.Scan(&verse)
		if err != nil {
			return models.SongLyrics{}, err
		}
		verses = append(verses, verse)
	}

	if len(verses) == 0 {
		return models.SongLyrics{}, fmt.Errorf("no lyrics found")
	}

	totalVerses := len(verses)

	page, err := strconv.Atoi(filters.Page)
	if err != nil || page < 1 {
		page = 1
	}

	if page > totalVerses {
		return models.SongLyrics{}, fmt.Errorf("page out of range")
	}

	return models.SongLyrics{
		ID:          filters.SongID,
		Group:       filters.Group,
		Song:        filters.Song,
		Lyrics:      verses[page-1],
		Page:        page,
		TotalVerses: totalVerses,
		HasNext:     page < totalVerses,
		HasPrev:     page > 1,
		NextPage:    page + 1,
		PrevPage:    page - 1,
	}, nil
}
