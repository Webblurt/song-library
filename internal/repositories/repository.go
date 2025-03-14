package repositories

import (
	"context"
	"errors"
	"fmt"
	models "song-library/internal/models"
	utils "song-library/internal/utils"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryInterface interface {
	Create(ctx context.Context, entity models.Entity) error
	Update(ctx context.Context, entity models.Entity) error
	Delete(ctx context.Context, entity models.Entity) error
	GetLibrary(ctx context.Context, filters models.LibraryFilters) (models.Library, error)
	GetSong(ctx context.Context, filters models.SongFilters) (models.SongLyrics, error)
}

type Repository struct {
	DB  *pgxpool.Pool
	log *utils.Logger
}

func NewRepository(cfg *utils.Config, log *utils.Logger) (*Repository, error) {
	tempRepo := &Repository{log: log}
	conn, err := tempRepo.CreateConnection(cfg)
	if err != nil {
		return nil, err
	}
	return &Repository{
		DB:  conn,
		log: log,
	}, nil
}

func (r *Repository) Create(ctx context.Context, entity models.Entity) error {
	columns, values, placeholders := buildInsertQuery(entity)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", entity.EntityName, columns, placeholders)
	_, err := r.DB.Exec(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, entity models.Entity) error {
	setClause, values := buildUpdateQuery(entity)
	if setClause == "" {
		return errors.New("no fields to update")
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", entity.EntityName, setClause, len(values)+1)
	values = append(values, entity.StringParameters["id"])

	_, err := r.DB.Exec(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, entity models.Entity) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", entity.EntityName)
	_, err := r.DB.Exec(ctx, query, entity.StringParameters["id"])
	if err != nil {
		return err
	}
	return nil
}

func buildInsertQuery(entity models.Entity) (columns string, values []interface{}, placeholders string) {
	var cols []string
	var ph []string
	var vals []interface{}

	i := 1
	for key, val := range entity.StringParameters {
		cols = append(cols, key)
		ph = append(ph, fmt.Sprintf("$%d", i))
		vals = append(vals, val)
		i++
	}
	for key, val := range entity.IntegerParameters {
		cols = append(cols, key)
		ph = append(ph, fmt.Sprintf("$%d", i))
		vals = append(vals, val)
		i++
	}
	for key, val := range entity.TimeParameters {
		cols = append(cols, key)
		ph = append(ph, fmt.Sprintf("$%d", i))
		vals = append(vals, val)
		i++
	}

	return strings.Join(cols, ", "), vals, strings.Join(ph, ", ")
}

func buildUpdateQuery(entity models.Entity) (setClause string, values []interface{}) {
	var setParts []string
	i := 1

	for key, val := range entity.StringParameters {
		if val == "" {
			continue
		}
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		values = append(values, val)
		i++
	}
	for key, val := range entity.IntegerParameters {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		values = append(values, val)
		i++
	}
	for key, val := range entity.TimeParameters {
		if val.IsZero() {
			continue
		}
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		values = append(values, val)
		i++
	}
	if len(setParts) == 0 {
		return "", nil
	}

	return strings.Join(setParts, ", "), values
}
