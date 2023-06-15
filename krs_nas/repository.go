package krs_nas

import (
	"gorm.io/gorm"
)

type Repository interface {
	// FindByID(nik string) (Stunting, error)
	// FindByPusHamil(nik string) (Stunting, error)
	// FindByStunting(nik string) (Stunting, error)
	// FindByBaduta(nik string) (Stunting, error)

	KrsByNas() ([]KrsNas, error)
	// FindByKec(ProvinsiID int, KabupatenID int, KecamatanID int) ([]Krs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) KrsByNas() ([]KrsNas, error) {
	var krsnases []KrsNas
	// err := r.db.Find(&faskeses, ProvinsiID).Error
	err := r.db.Find(&krsnases).Error
	return krsnases, err
}
