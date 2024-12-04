package krs_kecamatan

import (
	"gorm.io/gorm"
)

type Repository interface {
	KrsAllByKec(periode int) ([]KrsKec, error)
	// KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error)
	KrsByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKec, error)
	KrsKecById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Get All KRS dengan rincian by Kelurahan
func (r *repository) KrsAllByKec(periode int) ([]KrsKec, error) {
	var krsesKec []KrsKec
	err := r.db.Where("\"periode\" = ? ", periode).Find(&krsesKec).Error
	return krsesKec, err
}

// Get KRS by Kecamatan
func (r *repository) KrsByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKec, error) {
	var krsesKec []KrsKec
	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, Periode).Find(&krsesKec).Error
	return krsesKec, err
}

// Get KRS by Kecamatan
// func (r *repository) KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKel, error) {
// 	var krsesKel []KrsKel
// 	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"kode_kecamatan\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, Periode).Find(&krsesKel).Error
// 	return krsesKel, err
// }

// Get Detail KRS by Kode Kelurahan Kemendagri
func (r *repository) KrsKecById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error) {
	var krsesKec []KrsKec
	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"kode_kecamatan\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, Periode).Find(&krsesKec).Error
	return krsesKec, err
}
