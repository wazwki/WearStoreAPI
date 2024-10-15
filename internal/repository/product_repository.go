package repository

import (
	"WearStoreAPI/internal/models"
	"database/sql"
	"encoding/json"
)

type ProductStorage interface {
	FindById(id string) (*models.Item, error)
	FindAll() ([]*models.Item, error)
	Create(i *models.Item) error
	Update(i *models.Item, id string) error
	Delete(id string) error
}

type ProductRepository struct {
	DataBase *sql.DB
}

func (repo *ProductRepository) FindById(id string) (*models.Item, error) {
	var item models.Item
	var descriptionJSON []byte
	row := repo.DataBase.QueryRow(`SELECT price, title, photo, description FROM products_table WHERE id=$1`, id)

	if err := row.Scan(&item.Price, &item.Title, &item.Photo, &descriptionJSON); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(descriptionJSON, &item.Description); err != nil {
		return nil, err
	}

	return &item, nil
}

func (repo *ProductRepository) FindAll() ([]*models.Item, error) {
	rows, err := repo.DataBase.Query(`SELECT price, title, photo, description FROM products_table`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*models.Item{}

	if rows.Next() {
		var descriptionJSON []byte
		var item models.Item
		if err = rows.Scan(&item.Price, &item.Title, &item.Photo, &descriptionJSON); err != nil {
			return nil, err
		}
		if err = json.Unmarshal(descriptionJSON, &item.Description); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

func (repo *ProductRepository) Create(i *models.Item) error {
	descriptionJSON, err := json.Marshal(i.Description)
	if err != nil {
		return err
	}
	_, err = repo.DataBase.Exec(`INSERT INTO products_table(price, title, photo, description) VALUES ($1, $2, $3, $4)`, i.Price, i.Title, i.Photo, descriptionJSON)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) Update(i *models.Item, id string) error {
	descriptionJSON, err := json.Marshal(i.Description)
	if err != nil {
		return err
	}
	_, err = repo.DataBase.Exec(`UPDATE products_table SET price=$1, title=$2, photo=$3, description=$4 WHERE id=$5`, i.Price, i.Title, i.Photo, descriptionJSON, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) Delete(id string) error {
	_, err := repo.DataBase.Query(`DELETE FROM products_table WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
