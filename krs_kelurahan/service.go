package krs_kelurahan

type Service interface {
	KrsByProv(KodeDepdagriProvinsi string, Periode int) ([]KrsKel, error)
	KrsByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKel, error)
	KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKel, error)
	KrsKelById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, Periode int) ([]KrsKel, error)
	KrsAllByKel(periode int) ([]KrsKel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) KrsByProv(KodeDepdagriProvinsi string, Periode int) ([]KrsKel, error) {
	krses, err := s.repository.KrsByProv(KodeDepdagriProvinsi, Periode)
	return krses, err
}

func (s *service) KrsByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKel, error) {
	krses, err := s.repository.KrsByKab(KodeDepdagriProvinsi, KodeDepdagriKabupaten, Periode)
	return krses, err
}

func (s *service) KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKel, error) {
	krses, err := s.repository.KrsByKec(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, Periode)
	return krses, err
}

func (s *service) KrsAllByKel(periode int) ([]KrsKel, error) {
	krses, err := s.repository.KrsAllByKel(periode)
	return krses, err
}

func (s *service) KrsKelById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, Periode int) ([]KrsKel, error) {
	krses, err := s.repository.KrsKelById(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan, Periode)
	return krses, err
}
