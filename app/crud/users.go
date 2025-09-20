package crud

import (
	"app/models"
	repository "app/repositories"
)

func CreateUser(user *models.User) interface{} {
	repo := repository.UserRepo{}
	return repo.Create(user)
}

func GetUsers() ([]models.User, error) {
	repo := repository.UserRepo{}
	return repo.FindMany()
}
