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
	KrsByProvDetail(KodeDepdagriProvinsi string) ([]KrsNas, error)
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

func (r *repository) KrsByProvDetail(KodeDepdagriProvinsi string) ([]KrsNas, error) {
	var krsprovs []KrsNas
	// err := r.db.Find(&faskeses, ProvinsiID).Error
	err := r.db.Where("\"kode_depdagri_provinsi\" = ?  ", KodeDepdagriProvinsi).Find(&krsprovs).Error
	return krsprovs, err
}
