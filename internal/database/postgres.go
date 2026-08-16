package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/p-pannawit/salung-api/internal/config"
)

func Open(ctx context.Context, cfg config.DatabaseConfig) (*gorm.DB, *sql.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(connectionURL(cfg)), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("open gorm connection: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("get sql database: %w", err)
	}

	pingContext, cancel := context.WithTimeout(ctx, cfg.ConnectTimeout)
	defer cancel()
	if err := sqlDB.PingContext(pingContext); err != nil {
		sqlDB.Close()
		return nil, nil, fmt.Errorf("ping postgres: %w", err)
	}

	return gormDB, sqlDB, nil
}

func connectionURL(cfg config.DatabaseConfig) string {
	connectionURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Path:   cfg.Name,
	}

	query := connectionURL.Query()
	query.Set("sslmode", cfg.SSLMode)
	connectionURL.RawQuery = query.Encode()

	return connectionURL.String()
}
