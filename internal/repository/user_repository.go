package repository

import (
	"database/sql"
)

type UserRepository interface {
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(client *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: client,
	}
}
