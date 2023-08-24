package krs_kec

import (
	"gorm.io/gorm"
)

type Repository interface {
	// FindByID(nik string) (Stunting, error)
	// FindByPusHamil(nik string) (Stunting, error)
	// FindByStunting(nik string) (Stunting, error)
	// FindByBaduta(nik string) (Stunting, error)
	KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krs, error)
	KrsByKelDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) ([]Krs, error)
	// FindByKec(ProvinsiID int, KabupatenID int, KecamatanID int) ([]Krs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krs, error) {
	var krses []Krs
	// err := r.db.Find(&faskeses, ProvinsiID).Error
	err := r.db.Where("\"kode_depdagri_provinsi\" = ? AND \"kode_depdagri_kabupaten\" = ? AND \"kode_depdagri_kecamatan\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan).Find(&krses).Error
	return krses, err
}

func (r *repository) KrsByKelDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) ([]Krs, error) {
	var krses []Krs
	// err := r.db.Find(&faskeses, ProvinsiID).Error
	err := r.db.Where("\"kode_depdagri_provinsi\" = ? AND \"kode_depdagri_kabupaten\" = ? AND \"kode_depdagri_kecamatan\" = ? AND \"kode_depdagri_kelurahan\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan).Find(&krses).Error
	return krses, err
}
