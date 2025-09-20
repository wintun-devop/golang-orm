package repository

import (
	"app/config"
	"app/models"
)

type UserRepo struct{}

func (UserRepo) Create(user *models.User) interface{} {
	return config.DB.Create(user)
}

func (UserRepo) Update(id string, updates map[string]interface{}) error {
	return config.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

func (UserRepo) Remove(id string) error {
	return config.DB.Delete(&models.User{}, "id = ?", id).Error
}

func (UserRepo) SelectByID(id string) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, "id = ?", id).Error
	return &user, err
}

func (UserRepo) SelectByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, "email = ?", email).Error
	return &user, err
}

func (UserRepo) SelectByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, "username = ?", username).Error
	return &user, err
}

func (UserRepo) SelectByUniqueItems(email, username string) ([]models.User, error) {
	var users []models.User
	err := config.DB.Where("email = ? OR username = ?", email, username).Find(&users).Error
	return users, err
}

func (UserRepo) FindMany() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}
