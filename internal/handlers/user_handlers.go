package handlers

import (
	"WearStoreAPI/internal/service"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (u *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}
