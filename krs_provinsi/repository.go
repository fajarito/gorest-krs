package krs_provinsi

import (
	"gorm.io/gorm"
)

type Repository interface {
	KrsAllByProv(periode int) ([]KrsProv, error)
	// KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error)
	KrsByNas(Periode int) ([]KrsProv, error)
	KrsProvById(KodeDepdagriProvinsi string, Periode int) ([]KrsProv, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Get All KRS dengan rincian by Kelurahan
func (r *repository) KrsAllByProv(periode int) ([]KrsProv, error) {
	var krsesProv []KrsProv
	err := r.db.Where("\"periode\" = ? ", periode).Find(&krsesProv).Error
	return krsesProv, err
}

// Get KRS by Kecamatan
func (r *repository) KrsByNas(Periode int) ([]KrsProv, error) {
	var krsesProv []KrsProv
	err := r.db.Where("\"periode\" = ?", Periode).Find(&krsesProv).Error
	return krsesProv, err
}

// Get KRS by Kecamatan
// func (r *repository) KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKel, error) {
// 	var krsesKel []KrsKel
// 	err := r.db.Where("\"kode_provinsi\" = ? AND \"kode_kabupaten\" = ? AND \"kode_kecamatan\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, Periode).Find(&krsesKel).Error
// 	return krsesKel, err
// }

// Get Detail KRS by Kode Kelurahan Kemendagri
func (r *repository) KrsProvById(KodeDepdagriProvinsi string, Periode int) ([]KrsProv, error) {
	var krsesProv []KrsProv
	err := r.db.Where("\"kode_provinsi\" = ? AND \"periode\" = ?", KodeDepdagriProvinsi, Periode).Find(&krsesProv).Error
	return krsesProv, err
}
