package service

import (
	"WearStoreAPI/internal/repository"
)

type UserService struct {
	repo repository.UserStorage
}

func NewUserService(repo repository.UserStorage) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUser() {

}

func (u *UserService) CreateUser() {

}

func (u *UserService) UpdateUser() {

}

func (u *UserService) DeleteUser() {

}
