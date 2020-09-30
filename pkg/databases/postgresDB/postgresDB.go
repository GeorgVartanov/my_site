package postgresDB

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type DB struct {
	PostgresDB *sqlx.DB
}

func NewDB() (*DB, error) {
	db, err := sqlx.Open("postgres", os.Getenv("Url"))
	if err != nil {
		return nil,err
	}
	if err = db.Ping(); err != nil {
		return nil,err
	}
	return &DB{db}, nil
}



