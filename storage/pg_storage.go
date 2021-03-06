package storage

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// Storage impl
type pgStorage struct {
	db *sqlx.DB
}

func (s *pgStorage) QueryRow(query string, params ...interface{}) *sql.Row {
	return s.db.QueryRow(query, params...)
}

func (s *pgStorage) Exec(query string, params ...interface{}) error {
	_, err := s.db.Exec(query, params...)
	return err
}

func (s *pgStorage) QueryOne(dest interface{}, query string, params ...interface{}) error {
	return s.db.Get(dest, query, params...)
}

func (s *pgStorage) QueryAll(dest interface{}, query string, params ...interface{}) error {
	return s.db.Select(dest, query, params...)
}

func (s *pgStorage) Close() error {
	return s.db.Close()
}

func NewPgxStorage(dsn string, maxIdleConns, maxOpenConns int, connMaxLifetime time.Duration) (Storage, error) {
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// Prefer simple protocol means disabling implicit prepared statements
	config.PreferSimpleProtocol = true

	db, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(config))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	return &pgStorage{
		db,
	}, nil
}
