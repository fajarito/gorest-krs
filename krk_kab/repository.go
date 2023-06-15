package krk_kab

import (
	"gorm.io/gorm"
)

type Repository interface {
	// FindByID(nik string) (Stunting, error)
	// FindByPusHamil(nik string) (Stunting, error)
	// FindByStunting(nik string) (Stunting, error)
	// FindByBaduta(nik string) (Stunting, error)
	//KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krk, error)
	KrkByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, page int, pageSize int) ([]Krk, error)
	GetTotalKrkCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) (int64, error)
	// FindByKec(ProvinsiID int, KabupatenID int, KecamatanID int) ([]Krs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) KrkByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, page int, pageSize int) ([]Krk, error) {
	var krkes []Krk

	// err := r.db.Find(&faskeses, ProvinsiID).Error
	offset := (page - 1) * pageSize
	err := r.db.Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? ", KodeDepdagriProvinsi, KodeDepdagriKabupaten).Offset(offset).Limit(pageSize).Find(&krkes).Error

	if err != nil {
		return nil, err
	}

	return krkes, err
}

func (r *repository) GetTotalKrkCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) (int64, error) {
	var count int64
	if err := r.db.Model(&Krk{}).Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? ", KodeDepdagriProvinsi, KodeDepdagriKabupaten).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
