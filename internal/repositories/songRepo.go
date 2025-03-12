package repositories

import (
	"context"
	"fmt"
	models "song-library/internal/models"
	"strconv"
	"strings"
)

func (r *Repository) GetSong(ctx context.Context, filters models.SongFilters) (models.SongLyrics, error) {
	query := `SELECT lyrics FROM songs WHERE 1=1`
	args := []interface{}{}
	argID := 1

	if filters.SongID != "" {
		query += fmt.Sprintf(" AND id = $%d", argID)
		args = append(args, filters.SongID)
		argID++
	} else {
		if filters.Group != "" {
			query += fmt.Sprintf(" AND group_name ILIKE $%d", argID)
			args = append(args, "%"+filters.Group+"%")
			argID++
		}
		if filters.Song != "" {
			query += fmt.Sprintf(" AND song_name ILIKE $%d", argID)
			args = append(args, "%"+filters.Song+"%")
			argID++
		}
	}

	r.log.Debug("Executing query: ", query)

	var lyrics string
	err := r.DB.QueryRow(ctx, query, args...).Scan(&lyrics)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return models.SongLyrics{}, err
		}
		return models.SongLyrics{}, err
	}

	verses := strings.Split(lyrics, "\n\n")
	totalVerses := len(verses)

	page, err := strconv.Atoi(filters.Page)
	if err != nil || page < 1 {
		page = 1
	}

	if page > totalVerses {
		return models.SongLyrics{}, err
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
