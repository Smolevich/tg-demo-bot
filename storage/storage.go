package storage

import "database/sql"

type Storage interface {
	QueryRow(query string, params ...interface{}) *sql.Row
	Exec(query string, params ...interface{}) error
	QueryOne(dest interface{}, query string, params ...interface{}) error
	QueryAll(dest interface{}, query string, params ...interface{}) error
	Close() error
}
