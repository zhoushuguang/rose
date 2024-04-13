package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DSN         string        `yaml:"dsn"`
	Active      int           `yaml:"active"`
	Idle        int           `yaml:"idle"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

type DB struct {
	*sql.DB
}

func NewDB(conf *Config) *DB {
	if conf == nil {
		panic("conf cannot be nil")
	}
	d, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(conf.Active)
	d.SetMaxIdleConns(conf.Idle)
	d.SetConnMaxIdleTime(conf.IdleTimeout)

	return &DB{
		DB: d,
	}
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.QueryContext(ctx, query, args...)
}

func (db *DB) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.QueryRowContext(ctx, query, args...)
}

func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.ExecContext(ctx, query, args...)
}
