package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ShowUsers() ([]models.Users, error)
	GetUserByID(ID int) (models.Users, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowUsers() ([]models.Users, error) {
	var user []models.Users
	err := r.db.Find(&user).Error

	return user, err
}

func (r *repository) GetUserByID(ID int) (models.Users, error) {
	var user models.Users
	err := r.db.First(&user, ID).Error

	return user, err
}
