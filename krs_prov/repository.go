package krs_prov

import (
	"gorm.io/gorm"
)

type Repository interface {
	// FindByID(nik string) (Stunting, error)
	// FindByPusHamil(nik string) (Stunting, error)
	// FindByStunting(nik string) (Stunting, error)
	// FindByBaduta(nik string) (Stunting, error)

	KrsByProv(KodeDepdagriProvinsi string) ([]KrsProv, error)
	KrsByKabDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) ([]KrsProv, error)
	// FindByKec(ProvinsiID int, KabupatenID int, KecamatanID int) ([]Krs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) KrsByProv(KodeDepdagriProvinsi string) ([]KrsProv, error) {
	var krsprovs []KrsProv
	// err := r.db.Find(&faskeses, ProvinsiID).Error
	err := r.db.Where("\"kode_depdagri_provinsi\" = ?  ", KodeDepdagriProvinsi).Find(&krsprovs).Error
	return krsprovs, err
}

func (r *repository) KrsByKabDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) ([]KrsProv, error) {
	var krsprovs []KrsProv
	// err := r.db.Find(&faskeses, ProvinsiID).Error
	err := r.db.Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? ", KodeDepdagriProvinsi, KodeDepdagriKabupaten).Find(&krsprovs).Error
	return krsprovs, err
}
