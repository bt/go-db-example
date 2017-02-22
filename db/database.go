package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DbConfig interface {
	Dialect() string
}

type Db struct {
	DB *gorm.DB
}

func OpenDb(dbConfig DbConfig) (*Db, error) {
	switch dbConfig.Dialect() {
	case "sqlite3":
		return openSqliteDb(dbConfig.(*Sqlite3Config))
	default:
		return nil, fmt.Errorf("Unknown database dialect: '%s'", dbConfig.Dialect())
	}
}

func openSqliteDb(conf *Sqlite3Config) (*Db, error) {
	db, err := gorm.Open("sqlite3", conf.Path())
	if err != nil {
		return nil, err
	}
	return &Db{DB: db}, nil
}

// Closes the database.
func (db *Db) Close() error {
	return db.DB.Close()
}
