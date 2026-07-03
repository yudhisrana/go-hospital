package relational

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yudhisrana/go-hospital/internal/infra/config"
)

type DatabaseService struct {
	db *sql.DB
}

func NewDatabaseService(cfg config.DBConfig) (*DatabaseService, error) {
	// Open database connection
	db, err := sql.Open(cfg.DBDriver, cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("gagal terkoneksi ke database: %w", err)
	}

	db.SetMaxOpenConns(25)                 // Maximum open connections
	db.SetMaxIdleConns(5)                  // Maximum idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection
	db.SetConnMaxIdleTime(2 * time.Minute) // Maximum idle time of a connection

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("gagal ping ke database: %w", err)
	}

	return &DatabaseService{db: db}, nil
}

func (ds *DatabaseService) DB() *sql.DB {
	return ds.db
}

func (ds *DatabaseService) Close() error {
	if ds.db != nil {
		return ds.db.Close()
	}
	return nil
}
