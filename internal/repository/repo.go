package repository

import "WearStoreAPI/internal/models"

func GetWear(id string) (models.Item, error) {
	return models.Item{}, nil
}

func GetAllWear() ([]models.Item, error) {
	return nil, nil
}

func PostWear(i models.Item) error {
	return nil
}

func PatchWear(i models.Item) error {
	return nil
}

func DeleteWear(id string) error {
	return nil
}
