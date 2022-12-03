package stunting

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(nik string) (Stunting, error)
	FindByPusHamil(nik string) (Stunting, error)
	FindByStunting(nik string) (Stunting, error)
	FindByBaduta(nik string) (Stunting, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(nik string) (Stunting, error) {
	var stunting Stunting
	err := r.db.First(&stunting, nik).Error

	return stunting, err
}

func (r *repository) FindByStunting(nik string) (Stunting, error) {
	var stunting Stunting
	err := r.db.Where("nik = ? AND risiko_stunting = 'V'", nik).First(&stunting).Error
	return stunting, err
}

func (r *repository) FindByPusHamil(nik string) (Stunting, error) {
	var stunting Stunting
	err := r.db.Where("nik = ? AND pus_hamil = 'V'", nik).First(&stunting).Error
	return stunting, err
}

func (r *repository) FindByBaduta(nik string) (Stunting, error) {
	var stunting Stunting
	err := r.db.Where("nik = ? AND baduta = 'V'", nik).First(&stunting).Error
	return stunting, err
}
