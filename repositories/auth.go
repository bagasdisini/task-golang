package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterSV(user models.Users) (models.Users, error)
	RegisterUser(sv models.Users) (models.Users, error)
	Login(email string) (models.Users, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RegisterSV(user models.Users) (models.Users, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) RegisterUser(sv models.Users) (models.Users, error) {
	err := r.db.Preload("Users").Create(&sv).Error

	return sv, err
}

func (r *repository) Login(email string) (models.Users, error) {
	var users models.Users
	err := r.db.First(&users, "email=?", email).Error

	return users, err
}
