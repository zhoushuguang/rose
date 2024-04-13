package db

import (
	"github.com/zhoushuguang/rose/common/database"
)

type DB struct {
	db *database.DB
}

func NewDB(conf *database.Config) *DB {
	return &DB{
		db: database.NewDB(conf),
	}
}
