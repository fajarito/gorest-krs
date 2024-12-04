package krs_kabupaten

import (
	"gorm.io/gorm"
)

type Repository interface {
	KrsAllByKab(periode int) ([]KrsKab, error)
	// KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error)
	KrsByProv(KodeDepdagriProvinsi string, Periode int) ([]KrsKab, error)
	KrsKabById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKab, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Get All KRS dengan rincian by Kelurahan
func (r *repository) KrsAllByKab(periode int) ([]KrsKab, error) {
	var krsesKab []KrsKab
	err := r.db.Where("\"periode\" = ? ", periode).Find(&krsesKab).Error
	return krsesKab, err
}

// Get KRS by Kecamatan
func (r *repository) KrsByProv(KodeDepdagriProvinsi string, Periode int) ([]KrsKab, error) {
	var krsesKab []KrsKab
	err := r.db.Where("\"kode_provinsi\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, Periode).Find(&krsesKab).Error
	return krsesKab, err
}

// Get KRS by Kecamatan
// func (r *repository) KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKel, error) {
// 	var krsesKel []KrsKel
// 	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"kode_kecamatan\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, Periode).Find(&krsesKel).Error
// 	return krsesKel, err
// }

// Get Detail KRS by Kode Kelurahan Kemendagri
func (r *repository) KrsKabById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKab, error) {
	var krsesKab []KrsKab
	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, Periode).Find(&krsesKab).Error
	return krsesKab, err
}
