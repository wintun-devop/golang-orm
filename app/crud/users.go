package crud

import (
	"app/models"
	repository "app/repositories"
)

func CreateUser(user *models.User) error {
	repo := repository.UserRepo{}
	return repo.Create(user)
}
