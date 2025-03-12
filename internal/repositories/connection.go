package repositories

import (
	"context"
	"fmt"
	utils "song-library/internal/utils"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (r *Repository) CreateConnection(cfg *utils.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	r.log.Debug("Initializing database connection with DSN")

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = 10
	poolConfig.HealthCheckPeriod = 30 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.log.Debug("Creating database connection pool...")
	dbPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	r.log.Debug("Pinging database to verify connection...")
	if err := dbPool.Ping(ctx); err != nil {
		dbPool.Close()
		return nil, err
	}

	r.log.Debug("Database connection successfully established")
	return dbPool, nil
}
