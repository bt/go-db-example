package repository

import "github.com/bt/go-db-example/models"

type UserRepository interface {
	GetById(id int) (*models.User, error)
}
