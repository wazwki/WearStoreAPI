package repository

import (
	"WearStoreAPI/internal/models"
	"database/sql"
)

type UserStorage interface {
	FindById(id string) (*models.User, error)
	Create(i *models.User) error
	Update(i *models.User, id string) error
	Delete(id string) error
}

type UserRepository struct {
	DataBase *sql.DB
}

func (repo *UserRepository) FindById(id string) (*models.User, error) {
	return nil, nil
}

func (repo *UserRepository) Create(i *models.User) error {
	return nil
}

func (repo *UserRepository) Update(i *models.User, id string) error {
	return nil
}

func (repo *UserRepository) Delete(id string) error {
	return nil
}
