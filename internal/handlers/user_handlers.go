package handlers

import (
	"WearStoreAPI/internal/models"
	"WearStoreAPI/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (u *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	password := r.FormValue("password")

	token, err := u.service.LoginUser(id, password)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(72 * time.Hour),
	})
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := u.service.CreateUser(user)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	user, err := u.service.GetUser(id, tokenStr)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (u *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value
	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := u.service.UpdateUser(id, user, tokenStr); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	if err := u.service.DeleteUser(id, tokenStr); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
