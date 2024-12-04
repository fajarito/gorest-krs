package krs_kelurahan

import (
	"gorm.io/gorm"
)

type Repository interface {
	KrsAllByKel(periode int) ([]KrsKel, error)
	KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKel, error)
	KrsByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKel, error)
	KrsByProv(KodeDepdagriProvinsi string, Periode int) ([]KrsKel, error)
	KrsKelById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, Periode int) ([]KrsKel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Get All KRS dengan rincian by Kelurahan
func (r *repository) KrsAllByKel(periode int) ([]KrsKel, error) {
	var krsesKel []KrsKel
	err := r.db.Where("\"periode\" = ? ", periode).Find(&krsesKel).Error
	return krsesKel, err
}

// Get KRS by Kecamatan
func (r *repository) KrsByProv(KodeDepdagriProvinsi string, Periode int) ([]KrsKel, error) {
	var krsesKel []KrsKel
	err := r.db.Where("\"kode_provinsi\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, Periode).Find(&krsesKel).Error
	return krsesKel, err
}

// Get KRS by Kecamatan
func (r *repository) KrsByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKel, error) {
	var krsesKel []KrsKel
	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, Periode).Find(&krsesKel).Error
	return krsesKel, err
}

// Get KRS by Kecamatan
func (r *repository) KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKel, error) {
	var krsesKel []KrsKel
	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"kode_kecamatan\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, Periode).Find(&krsesKel).Error
	return krsesKel, err
}

// Get Detail KRS by Kode Kelurahan Kemendagri
func (r *repository) KrsKelById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, Periode int) ([]KrsKel, error) {
	var krsesKel []KrsKel
	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"kode_kecamatan\" = ? AND \"kode_kelurahan\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan, Periode).Find(&krsesKel).Error
	return krsesKel, err
}
