package krs_kecamatan

type Service interface {
	KrsByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKec, error)
	// KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error)
	KrsKecById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error)
	KrsAllByKec(periode int) ([]KrsKec, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) KrsByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKec, error) {
	krses, err := s.repository.KrsByKab(KodeDepdagriProvinsi, KodeDepdagriKabupaten, Periode)
	return krses, err
}

func (s *service) KrsAllByKec(periode int) ([]KrsKec, error) {
	krses, err := s.repository.KrsAllByKec(periode)
	return krses, err
}

func (s *service) KrsKecById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error) {
	krses, err := s.repository.KrsKecById(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, Periode)
	return krses, err
}
