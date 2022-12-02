package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type PresenceRepository interface {
	ShowPresences() ([]models.EPresence, error)
	GetPresenceByID(ID int) (models.EPresence, error)
	CreatePresence(Presence models.EPresence) (models.EPresence, error)
	UpdatePresence(Presence models.EPresence, ID int) (models.EPresence, error)
}

func RepositoryPresence(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowPresences() ([]models.EPresence, error) {
	var Presences []models.EPresence
	err := r.db.Find(&Presences).Error

	return Presences, err
}

func (r *repository) GetPresenceByID(ID int) (models.EPresence, error) {
	var Presence models.EPresence
	err := r.db.First(&Presence, ID).Error

	return Presence, err
}

func (r *repository) CreatePresence(Presence models.EPresence) (models.EPresence, error) {
	err := r.db.Create(&Presence).Error

	return Presence, err
}

func (r *repository) UpdatePresence(Presence models.EPresence, ID int) (models.EPresence, error) {
	err := r.db.Model(&Presence).Where("id=?", ID).Updates(&Presence).Error

	return Presence, err
}
