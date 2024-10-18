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
	var user models.User
	row := repo.DataBase.QueryRow(`SELECT email, first_name, last_name, password FROM users_table WHERE id=$1`, id)

	if err := row.Scan(&user.Email, &user.FirstName, &user.LastName, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Create(u *models.User) error {
	_, err := repo.DataBase.Exec(`INSERT INTO users_table(email, first_name, last_name, password) VALUES ($1, $2, $3, $4)`, u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Update(u *models.User, id string) error {
	_, err := repo.DataBase.Exec(`UPDATE users_table SET email=$1, first_name=$2, last_name=$3, password=$4 WHERE id=$5`, u.Email, u.FirstName, u.LastName, u.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(id string) error {
	_, err := repo.DataBase.Query(`DELETE FROM users_table WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
