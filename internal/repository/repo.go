package repository

import (
	"WearStoreAPI/internal/models"
	"database/sql"
)

type ProductStorage interface {
	FindById(id string) (*models.Item, error)
	FindAll() ([]*models.Item, error)
	Create(i *models.Item) error
	Patch(i *models.Item, id string) error
	Delete(id string) error
}

type ProductRepository struct {
	DataBase *sql.DB
}

func (repo *ProductRepository) FindById(id string) (*models.Item, error) {
	return &models.Item{}, nil
}

func (repo *ProductRepository) FindAll() ([]*models.Item, error) {
	return nil, nil
}

func (repo *ProductRepository) Create(i *models.Item) error {
	return nil
}

func (repo *ProductRepository) Patch(i *models.Item, id string) error {
	return nil
}

func (repo *ProductRepository) Delete(id string) error {
	return nil
}
