package krs_kec

type Service interface {
	KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krs, error)
	KrsByKelDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) ([]Krs, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krs, error) {
	faskeses, err := s.repository.KrsByKec(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan)
	return faskeses, err
}

func (s *service) KrsByKelDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) ([]Krs, error) {
	faskeses, err := s.repository.KrsByKelDetail(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan)
	return faskeses, err
}
