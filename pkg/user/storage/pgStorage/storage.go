package pgStorage

import "github.com/GeorgVartanov/my_site/pkg/databases/postgresDB"

type UserStorage struct {
	DB *postgresDB.DB
}

func NewUserStorage(DB *postgresDB.DB) *UserStorage {
	return &UserStorage{DB: DB}
}
