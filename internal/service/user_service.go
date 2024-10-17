package service

import (
	"WearStoreAPI/internal/models"
	"WearStoreAPI/internal/repository"
)

type UserService struct {
	repo repository.UserStorage
}

func NewUserService(repo repository.UserStorage) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) LoginUser(id string, password string) (string, error) {
	// сравнить хеш пароля с хешем пароля в таблице если совпал то выдаем токен
}

func (u *UserService) GetUser(id string, token string) (*models.User, error) {
	// считать из jwt почту если почта совпадает с почтой по данному id выполнить действия с записью пользователя
}

func (u *UserService) CreateUser(*models.User) error {
	// хеш пароль из user
}

func (u *UserService) UpdateUser(id string, user *models.User, token string) error {
	// хеш пароль из user
	// считать из jwt почту если почта совпадает с почтой по данному id выполнить действия с записью пользователя
}

func (u *UserService) DeleteUser(id string, token string) error {
	// считать из jwt почту если почта совпадает с почтой по данному id выполнить действия с записью пользователя
}
