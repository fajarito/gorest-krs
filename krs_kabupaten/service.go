package krs_kabupaten

type Service interface {
	KrsByProv(KodeDepdagriProvinsi string, Periode int) ([]KrsKab, error)
	// KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error)
	KrsKabById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKab, error)
	KrsAllByKab(periode int) ([]KrsKab, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) KrsByProv(KodeDepdagriProvinsi string, Periode int) ([]KrsKab, error) {
	krses, err := s.repository.KrsByProv(KodeDepdagriProvinsi, Periode)
	return krses, err
}

func (s *service) KrsAllByKab(periode int) ([]KrsKab, error) {
	krses, err := s.repository.KrsAllByKab(periode)
	return krses, err
}

func (s *service) KrsKabById(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, Periode int) ([]KrsKab, error) {
	krses, err := s.repository.KrsKabById(KodeDepdagriProvinsi, KodeDepdagriKabupaten, Periode)
	return krses, err
}
