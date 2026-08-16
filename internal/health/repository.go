package health

import (
	"context"
	"database/sql"
)

type Repository interface {
	Ping(context.Context) error
}

type PostgresRepository struct {
	database *sql.DB
}

func NewPostgresRepository(database *sql.DB) *PostgresRepository {
	return &PostgresRepository{database: database}
}

func (repository *PostgresRepository) Ping(ctx context.Context) error {
	return repository.database.PingContext(ctx)
}
