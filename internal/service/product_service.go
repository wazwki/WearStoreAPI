package service

import (
	"WearStoreAPI/internal/models"
	"WearStoreAPI/internal/repository"
)

type ProductService struct {
	repo repository.ProductStorage
}

func NewProductService(repo repository.ProductStorage) *ProductService {
	return &ProductService{repo: repo}
}

func (p *ProductService) GetWearData(id string) (*models.Item, error) {
	wear, err := p.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	return wear, nil
}

func (p *ProductService) GetAllWearData() ([]*models.Item, error) {
	wears, err := p.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return wears, nil
}

func (p *ProductService) CreateWear(i *models.Item) error {
	err := p.repo.Create(i)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) UpdateWear(i *models.Item, id string) error {
	err := p.repo.Update(i, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) DeleteWear(id string) error {
	err := p.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
