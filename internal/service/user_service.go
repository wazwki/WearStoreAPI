package service

import (
	"WearStoreAPI/internal/models"
	"WearStoreAPI/internal/repository"
	"WearStoreAPI/pkg/auth"
	"os"
)

type UserService struct {
	repo repository.UserStorage
}

func NewUserService(repo repository.UserStorage) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) LoginUser(id string, password string) (string, error) {
	user, err := u.repo.FindById(id)
	if err != nil {
		return "", err
	}

	if auth.ComparePassword(user.Password, password) {
		token, err := auth.CreateToken(user.Email, []byte(os.Getenv("JWT_KEY")))
		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", err
}

func (u *UserService) GetUser(id string, token string) (*models.User, error) {
	user, err := u.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	claims, err := auth.CheckToken(token, []byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	if useremail, ok := claims["email"].(string); ok {
		if user.Email == useremail {
			return user, nil
		}
	}

	return nil, err
}

func (u *UserService) CreateUser(user *models.User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = u.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) UpdateUser(id string, user *models.User, token string) error {
	userCheck, err := u.repo.FindById(id)
	if err != nil {
		return err
	}

	claims, err := auth.CheckToken(token, []byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return err
	}

	if useremail, ok := claims["email"].(string); ok {
		if userCheck.Email == useremail {

			hashedPassword, err := auth.HashPassword(user.Password)
			if err != nil {
				return err
			}
			user.Password = hashedPassword

			err = u.repo.Update(user, id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *UserService) DeleteUser(id string, token string) error {
	userCheck, err := u.repo.FindById(id)
	if err != nil {
		return err
	}

	claims, err := auth.CheckToken(token, []byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return err
	}

	if useremail, ok := claims["email"].(string); ok {
		if userCheck.Email == useremail {
			err := u.repo.Delete(id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
